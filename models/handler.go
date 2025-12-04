package models


import (
	"net/http"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type ChatRoomHandler struct {
	ChatRoomServ *ChatRoomServer
}

// Handler to hold the chat room server
func NewChatRoomHandler(ChH *ChatRoomServer) *ChatRoomHandler {
	return &ChatRoomHandler{
		ChatRoomServ: ChH,
	}
}


// for creating a room
func (ChH *ChatRoomHandler) CreateNewRoom(c *fiber.Ctx) error {
	// create a new room request
	// Change to return the error and the .JSON response

	var newRoomReq CreateRoomReq

	// validate the request from the user
	if err := c.BodyParser(&newRoomReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Optional: prevent duplicate rooms
	if _, exists := ChH.ChatRoomServ.Rooms[newRoomReq.ID]; exists {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": "room already exists",
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
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"ID":     newRoomReq.ID,
		"Name":   newRoomReq.Name,
		"status": http.StatusOK,
	})
}

// User decides to join a room

func (ChH *ChatRoomHandler) JoinRoom(c *websocket.Conn) {
	// (Optional) ensure the room exists
	if _, ok := ChH.ChatRoomServ.Rooms[c.Params("roomId")]; !ok {
		// you can auto-create, or just close the connection
		// here we'll just close
		_ = c.Close()
		return
	}

	newUser := &Client{
		Conn:     c,
		Message:  make(chan *ChatMessage, 10),
		ID:       c.Params("userId"),
		RoomID:   c.Params("roomId"),
		Username: c.Params("username"),
	}

	// system message: user joined
	newUserMessage := &ChatMessage{
		Content:  "A new user has joined the chat room",
		RoomID:   c.Params("roomId"),
		Username: c.Params("username"),
	}


	// register the new user
	ChH.ChatRoomServ.Register <- newUser
	ChH.ChatRoomServ.Broadcast <- newUserMessage

	// write the actual new user message to group after introduction
	go newUser.writeMessage()

	// start read loop  (blocks until disconnect)
	newUser.ReadMessage(ChH.ChatRoomServ)
}