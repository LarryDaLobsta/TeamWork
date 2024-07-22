package models

// This will house some of the code from the main server file that will model chatrooms, build chatrooms, delete chatrooms,
// update users in chatrooms, etc. All things chatrooms

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"log"
	_ "log"
	"net/http"
	_ "os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// websocket server structs
type ChatRoomServer struct {
	chatRoomServerName string
	broadcast          chan *ChatMessage
	Rooms              map[string]*ChatRoom
	Register           chan *Client
	Unregister         chan *Client
}

// This is the chat room struct for each chatroom per project
// Need to make sure each pkoject gets assigned one ChatRoom
// until I can make smaller chatrooms for individual pieces

type ChatRoom struct {
	ID        string
	Name      string
	Project   string
	ProjectId int
	clients   map[string]*Client
	broadcast chan *Message
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *ChatMessage
	ID       string `json:"id"`
	RoomID   string `json:"roomid"`
	Username string `json:"username"`
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

		c.Conn.WriteJSON(message)
	}
}

// read the messages from the hub/chat room
func (c *Client) ReadMessage(ChS ChatRoomServer) {
	defer func() {
		ChS.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf(" error: v%", err)
			}
			break
		}

		// handle the message if the web socket connection is still good

		msg := &ChatMessage{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		// now broadcast the message to the correct Room
		ChS.broadcast <- msg
	}
}

type ChatMessage struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ChatRoomHandler struct {
	ChatRoomServ *ChatRoomServer
}

// Handler to hold the chat room server
func NewChatRoomHandler(h *ChatRoomServer) *ChatRoomHandler {
	return &ChatRoomHandler{
		ChatRoomServ: h,
	}
}

// Creates a actual chat room server
func NewChatroomServer() *ChatRoomServer {
	// create a new chatroom ent framework to be saved in the database to be saved to the ent database
	return &ChatRoomServer{
		Rooms:      make(map[string]*ChatRoom),
		broadcast:  make(chan *ChatMessage, 5),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (ChS *ChatRoomServer) StartServer() {
	for {
		select {
		case c1 := <-ChS.Register:

			// check to make sure the room exists
			if _, ok := ChS.Rooms[c1.RoomID]; ok {
				room := ChS.Rooms[c1.RoomID]

				// check to make sure the user is not already in the room
				if _, ok := room.clients[c1.ID]; !ok {
					// add the user if he is not in the room already
					room.clients[c1.ID] = c1
				}

			}
		}

		// if we are registering

		// if we are unregistering

		// broadcast message to clients in a room
	}
}

// for creating a room
func (ChH *ChatRoomHandler) CreateNewRoom(ctx *fiber.Ctx) {
	// create a new room request

	var newRoomReq CreateRoomReq

	// validate the request from the user
	if err := ctx.BodyParser(newRoomReq); err != nil {
		ctx.JSON(fiber.Map{
			"ID":     newRoomReq.ID,
			"Name":   newRoomReq.Name,
			"status": http.StatusBadRequest,
		})

		return
	}

	// if a good request
	// add this created room to the server
	ChH.ChatRoomServ.Rooms[newRoomReq.ID] = &ChatRoom{
		ID:      newRoomReq.ID,
		Name:    newRoomReq.Name,
		clients: make(map[string]*Client),
	}

	ctx.JSON(fiber.Map{
		"ID":     newRoomReq.ID,
		"Name":   newRoomReq.Name,
		"status": http.StatusOK,
	})
}

// User decides to join a room

func (ChH *ChatRoomHandler) JoinRoom(ctx *fiber.Ctx, c *websocket.Conn) {
	// authentication will happen in the .Use route in the server
	// then after authentication we will move toward the next method in the stack
	// which will get the websocket route and connection which will be passed in here

	newUser := &Client{
		Conn:     c,
		Message:  make(chan *ChatMessage, 10),
		ID:       ctx.Params("userId"),
		RoomID:   ctx.Query("roomId"),
		Username: ctx.Params("username"),
	}

	// If succesfful create client, create messsages, register user to hub

	// create the messages to notify new user joining
	newUserMessage := &ChatMessage{
		Content:  "A new user has joined the chat room",
		RoomID:   ctx.Params("roomId"),
		Username: ctx.Query("username"),
	}

	// register the new user
	ChH.ChatRoomServ.Register <- newUser
	ChH.ChatRoomServ.broadcast <- newUserMessage

	// write the actual new user message to group after introduction
	go newUser.writeMessage()
	newUser.ReadMessage(*ChH.ChatRoomServ)
}

func StartServer()

// add users to a webserver/chatroom
// need to clean this up to be able to add Handle incoming and outgoing websocket connections
func (ChH *ChatRoomHandler) HandleWebSocket(ctx *websocket.Conn) {
	// we have a new websocket conn coming down the pipeline meaning a user wants to join a room

	// REgister a new Client
	s.clients[ctx] = true
	defer func() {
		delete(s.clients, ctx)
		ctx.Close()
	}()

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error", err)
			break
		}

		// send the message to the broadcast channel
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Fatalf("Error Umarshalling")
		}
		s.broadcast <- &message
	}
}

func (s *WebSocketServer) HandleMessages() {
	for {
		msg := <-s.broadcast

		// Send the message to all Clients

		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, getMessageTemplate(msg))
			if err != nil {
				log.Printf("Write Error: %v", err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}
