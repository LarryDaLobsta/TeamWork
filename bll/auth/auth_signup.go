package bll

import (
	"context"
	"teamplayer/ent"
	models "teamplayer/models"
	"fmt"
	dal "teamplayer/dal"
)

const maxFirstNameLength = 50
const maxLastNameLength  = 65
const maxEmailLength     = 254
const minUserNameLength  = 6
const maxUserNameLength  = 24
const minPasswordLength  = 12
//const maxPasswordLength  = 24

func UserSignUp(ctx context.Context, client *ent.Client, NewUser models.UserSignUp) error {

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
	if err := ValidateEmailEntry("email", "email", email , maxEmailLength); err != nil {
		return err
	}

	//valide username
	if err := ValidateUsernameEntry("user_name", "user name", userName, minUserNameLength, maxUserNameLength); err != nil {
		return err
	}

	//validate password
	if err := ValidatePasswordEntry("password", "password", password, minPasswordLength); err != nil {
		return err
	}

	hash, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("password hash unsuccessful: %w", err)
	}
	// if all goes well then call the function to go to the database
	// return from this should be if successful then return true and no error

	//dal object to protect database 

	SuccNewUser := models.UserRecord{
		FirstName:		firstName,
		LastName:		lastName,
		UserName:		userName,
		Email:			email,
		PasswordHash:	hash,
	}

	createdUserStatus := dal.CreateUser(SuccNewUser, client, ctx)
	if createdUserStatus != nil {
		return fmt.Errorf("User Creation unsuccessful: %w", createdUserStatus)
	}

	return nil

}