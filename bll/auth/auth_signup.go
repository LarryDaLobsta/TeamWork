package bll

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"github.com/gofiber/contrib/websocket"
	dal "teamplayer/dal"
)

const maxFirstNameLength = 50
const maxLastNameLength  = 65
const maxEmailLength     = 254
const minUserNameLength  = 6
const maxUserNameLength  = 24
const minPasswordLength  = 12
//const maxPasswordLength  = 24

func userSignUp(ctx context.Context, client *ent.Client, NewUser SignUpInput) error {

	firstName := NewUser.FirstName
	lastName  := NewUser.LastName	
	userName  := NewUser.UserName
	email 	  := NewUser.Email
	password  := NewUser.Password

	// validate first name
	if err := ValidateNameEntry("first_name", "first name", firstName, maxFirstNameLength); err != nil {
		return err
	}

	//validate last name
	if err := ValidateNameEntry("last_name", "last name", lastName, maxLastNameLength); err != nil {
		return err
	}

	//validate email
	if err := ValidateUsernameEntry("username", "username", userName , minUserNameLength, maxUserNameLength); err != nil {
		return error
	}

	//validate password
	if err := ValidatePasswordEntry(password); err != nil {
		return error
	}

	hash, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("password hash unsuccessful: %w", err)
	}
	// if all goes well then call the function to go to the database
	// return from this should be if successful then return true and no error

	dal.CreateUser()

}