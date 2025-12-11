package bll

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"github.com/gofiber/contrib/websocket"
	dal "teamplayer/dal"
)




func userSignUp()(*User, client *ent.Clinet, error){

	// validate user information
	if err := ValidateNameEntry( *User.FirstName); err != nil {
		return nil, error

	}
	if err := ValidateNameEntry( *User.FirstName); err != nil {
		return nil, error
	}
	if err := ValidateEmailEntry( *User.Email); err != nil {
		return nil, error
	}
	if err := ValidateUserNameEntry( *User.UserName); err != nil {
		return nil, error
	}
	if err := ValidatePasswordEntry( *User.Password); err != nil {
		return nil, error
	}


	// if all goes well then call the function to go to the database

	dal.CreateUser()

}