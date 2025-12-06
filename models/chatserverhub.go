package models

import (
	"fmt"
	"log"
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
        case client := <-ChS.Register:
           room, ok := ChS.Rooms[client.RoomID]
		   if !ok {
				log.Printf("Register: room %s not found", client.RoomID)
				continue
		   }
		   
		   if room.AddClient(client){
				log.Printf("Register: %s joined room %s", client.Username, client.RoomID)
		   } else{
				log.Printf("Register: %s already joined room %s", client.Username, client.RoomID)		
		   }
			
		   

        case client := <-ChS.Unregister:
			// make sure room exists
			room, ok := ChS.Rooms[client.RoomID]
			if !ok {
				log.Printf("Register: room %s not found", client.RoomID)
				continue
		   	}

			//make sure user exists in that room
			if !room.RemoveClient(client){
				log.Printf("Unregisterer: %s not found in room %s", client.Username, client.RoomID)
				continue
			}

			leaveMsg := &ChatMessage{
				Content:  fmt.Sprintf("User %s has left the chat room", client.Username),
				RoomID:   client.RoomID,
				Username: client.Username,
			}

			for _, other := range room.Clients {
				other.Message <- leaveMsg
			}
		
			// Optionally delete empty rooms
			if room.IsEmpty() {
				log.Printf("Room %s is empty, deleting...", room.ID)
				delete(ChS.Rooms, room.ID)
			}

        case m := <-ChS.Broadcast:
            if room, ok := ChS.Rooms[m.RoomID]; ok {
                log.Printf("Broadcast in %s: %s: %s", m.RoomID, m.Username, m.Content)
                for _, client := range room.Clients {
                    client.Message <- m
                }
            } else {
                log.Printf("Broadcast: room %s not found", m.RoomID)
            }
        }
    }
}