package models

// This will house some of the code from the main server file that will model chatrooms, build chatrooms, delete chatrooms,
// update users in chatrooms, etc. All things chatrooms


// This is the chat room struct for each chatroom per project
// Need to make sure each pkoject gets assigned one ChatRoom
// until I can make smaller chatrooms for individual pieces

type ChatRoom struct {
	ID        string
	Name      string
	Project   string
	ProjectId int
	Clients   map[string]*Client
}


package models


import (
	"net/http"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type ChatRoomHandler struct {
	ChatRoomServ *ChatRoomServer
}

func NewChatRoom(id, name string) *ChatRoom {
    return &ChatRoom{
        ID:      id,
        Name:    name,
        Clients: make(map[string]*Client),
    }
}

func (r *ChatRoom) AddClient(c *Client) bool 
{
	// check if the user exists in the room already
	if _, exists := r.Clients[c.ID]; exists {
		// already in the room
		return false
	}

	// add user if not there
	r.Clients[c.ID] = c

	return true
}

func (r *ChatRoom) RemoveClient(c *Client) bool 
{
	// check if the user exists in the room already
	if _, exists := r.Clients[c.ID]; !exists {
		// already removed
		return false
	}

	// delete from room
	delete(r.Clients, c.ID)

	// Close channel for the user removed
	close(c.Message)

	return true

}

func (r *ChatRoom) IsEmpty() bool 
{
	return len(r.Clients) == 0
}

