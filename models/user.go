package models

import (
	_ "fmt"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

type User struct {
	// adding to the ent struct for dealing with application post requests
	UUID_Id   uuid.UUID `json:"UUID"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
}
