package models

// This will house some of the code from the main server file that will model chatrooms, build chatrooms, delete chatrooms,
// update users in chatrooms, etc. All things chatrooms

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"log"
	_ "log"
	_ "os"

	"github.com/gofiber/contrib/websocket"
)

// websocket server structs
type WebSocketServer struct {
	chatRoomId   int
	chatRoomName string
	clients      map[*websocket.Conn]bool
	broadcast    chan *Message
}

func NewWebSocket() *WebSocketServer {
	// create a new chatroom ent framework to be saved in the database to be saved to the ent database
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
	}
}

// add users to a webserver/chatroom
func (s *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {
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
