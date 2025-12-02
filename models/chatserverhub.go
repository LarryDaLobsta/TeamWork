package models

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)



// websocket server structs
type ChatRoomServer struct {
	chatRoomServerName string
	Broadcast          chan *ChatMessage
	Rooms              map[string]*ChatRoom
	Register           chan *Client
	Unregister         chan *Client
}


//constructor to create a ChatRoomServer
func NewChatRoomServer() *ChatRoomServer{
	return &ChatRoomServer{
		Rooms:      make(map[string]*ChatRoom),
        Broadcast:  make(chan *ChatMessage, 5),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
	}
}


// Start Server
func (ChS *ChatRoomServer) StartServer() {
	log.Println("ChatRoomServer started")
	for {
        select {
        case c1 := <-ChS.Register:
            if room, ok := ChS.Rooms[c1.RoomID]; ok {
                if _, exists := room.Clients[c1.ID]; !exists {
                    log.Printf("Register: %s joined room %s", c1.Username, c1.RoomID)
                    room.Clients[c1.ID] = c1
                }
            } else {
                log.Printf("Register: room %s not found", c1.RoomID)
            }

        case c1 := <-ChS.Unregister:
            if room, ok := ChS.Rooms[c1.RoomID]; ok {
                if _, exists := room.Clients[c1.ID]; exists {
                    log.Printf("Unregister: %s leaving room %s", c1.Username, c1.RoomID)
                    delete(room.Clients, c1.ID)
                    close(c1.Message)
                }
            }

        case m := <-ChS.broadcast:
            if room, ok := ChS.Rooms[m.RoomID]; ok {
                log.Printf("Broadcast in %s: %s: %s", m.RoomID, m.Username, m.Content)
                for _, c1 := range room.Clients {
                    c1.Message <- m
                }
            } else {
                log.Printf("Broadcast: room %s not found", m.RoomID)
            }
        }
    }
}