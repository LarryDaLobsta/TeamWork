package models

// This will house some of the code from the main server file that will model chatrooms, build chatrooms, delete chatrooms,
// update users in chatrooms, etc. All things chatrooms

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)



// This is the chat room struct for each chatroom per project
// Need to make sure each pkoject gets assigned one ChatRoom
// until I can make smaller chatrooms for individual pieces

type ChatRoom struct {
	ID        string
	Name      string
	Project   string
	ProjectId int
	Clients   map[string]*Client
	broadcast chan *Message
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *ChatMessage
	ID       string `json:"id"`
	RoomID   string `json:"roomid"`
	Username string `json:"username"`
}

type ChatRoomHandler struct {
	ChatRoomServ *ChatRoomServer
}


// write the message to the message attibute for client
func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		// Build HTML snippet for HTMX to insert
		htmlMsg := fmt.Sprintf(
			`<div id="messages" hx-swap-oob="beforeend">
				<div class="message"><strong>%s:</strong> %s</div>
				</div>`,
			html.EscapeString(message.Username),
			html.EscapeString(message.Content),
		)

		if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(htmlMsg)); err != nil {
			log.Println("write message error:", err)
			return
		}
	}
}

// read the messages from the hub/chat room
func (c *Client) ReadMessage(ChS ChatRoomServer) {
	defer func() {
		ChS.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, raw, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("read message error: %v", err)
			}
			break
		}

		// handle the message if the web socket connection is still good
		log.Printf("Raw from client: %s", raw)

		// HTMX WebSocket payload format
		var payload struct {
			Text    string                 `json:"text"`
			Headers map[string]interface{} `json:"HEADERS"`
		}

		if err := json.Unmarshal(raw, &payload); err != nil {
			log.Printf("json error: %v", err)
			continue
		}

		log.Printf("Parsed text from %s: %s", c.Username, payload.Text)

		msg := &ChatMessage{
			Content:  payload.Text,
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		// now broadcast the message to the correct Room
		ChS.broadcast <- msg
	}
}




// Handler to hold the chat room server
func NewChatRoomHandler(h *ChatRoomServer) *ChatRoomHandler {
	return &ChatRoomHandler{
		ChatRoomServ: h,
	}
}




// for creating a room
func (ChH *ChatRoomHandler) CreateNewRoom(c *fiber.Ctx) error {
	// create a new room request
	// Change to return the error and the .JSON response

	var newRoomReq CreateRoomReq

	// validate the request from the user
	if err := c.BodyParser(&newRoomReq); err != nil {
		return c.JSON(fiber.Map{
			"ID":     newRoomReq.ID,
			"Name":   newRoomReq.Name,
			"status": http.StatusBadRequest,
		})
	}

	// if a good request
	// add this created room to the server
	ChH.ChatRoomServ.Rooms[newRoomReq.ID] = &ChatRoom{
		ID:      newRoomReq.ID,
		Name:    newRoomReq.Name,
		Clients: make(map[string]*Client),
	}

	c.Locals("allowed", true)
	return c.JSON(fiber.Map{
		"ID":     newRoomReq.ID,
		"Name":   newRoomReq.Name,
		"status": http.StatusOK,
	})
}

// User decides to join a room

func (ChH *ChatRoomHandler) JoinRoom(c *websocket.Conn) {
	// authentication will happen in the .Use route in the server
	// then after authentication we will move toward the next method in the stack
	// which will get the websocket route and connection which will be passed in here

	newUser := &Client{
		Conn:     c,
		Message:  make(chan *ChatMessage, 10),
		ID:       c.Params("userId"),
		RoomID:   c.Params("roomId"),
		Username: c.Params("username"),
	}

	// register the new user
	ChH.ChatRoomServ.Register <- newUser
	ChH.ChatRoomServ.broadcast <- newUserMessage

	// write the actual new user message to group after introduction
	go newUser.writeMessage()
	newUser.ReadMessage(*ChH.ChatRoomServ)
}
