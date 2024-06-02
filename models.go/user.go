//This will hold the user model, these will some of the main to go into the database 
// user: CRUD ability for to do items, messaging ability, CRUD for chat rooms

type User struct {
	UserID  int
	//UserUUID
	FirstName string
	LastName string
	UserName string

	//ModUUID
	//GroupchatUUID

	//may need to add the messages that are associated with the  user
	//may need to add the items that a user has created/responsible
	// chat rooms that the user has
}


type Message struct {
	MessageID int
	//MessageUUID
	//who sent the message
	SenderUUID
	// what group chat does the message belong to
	GroupUUID
	// when was the message sent 
	TimeSent time
	// when was the message received
	TimeRecv time
}

type GroupChat struct {
	GroupchatID int
	//GroupchateUUID 
	// creater of the groupchat
	GroupchatCreaterUUID
	// if you are a mod for this groupchat you will have
	// a specific UUID for each groupchat
	GroupchatModUUID
}



