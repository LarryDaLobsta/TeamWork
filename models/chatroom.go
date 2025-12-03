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


