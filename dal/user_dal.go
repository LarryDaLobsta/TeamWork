package dal

import (
	"context"
	"fmt"
	"log"
	"teamplayer/ent"
	"teamplayer/ent/user"
	M "teamplayer/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// Creating a user
// Developer note need to check the new addition of a user
//
//	Check to see if a user has the same user name and password
func CreateUser(ctx context.Context, c *fiber.Ctx, client *ent.Client) error {
	newUser := new(M.User)
	// grab and put in struct
	// return false if not able to do it
	if err := c.BodyParser(newUser); err != nil {
		return err
	}

	err := client.User.
		Create().
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		SetUsername(newUser.UserName).
		SetPassword(newUser.Password).
		SaveX(ctx) // See if different type of save occurs do to panic
		// check to make sure the save is successful
	if err != nil {
		return fmt.Errorf("Failed creating the new user: %w", err)
	}
	return nil
}

// returns true if user found or false if user not found
func CheckUser(ctx context.Context, c *fiber.Ctx, client *ent.Client) (bool, error) {
	// Checks to see if a user already with a password or username provided by a new user

	// create a new user struct
	newUser := new(M.User)

	// grab and put in struct
	// return false if not able to do it
	if err := c.BodyParser(newUser); err != nil {
		return true, err
	}

	// check to see if the username and password is in the database

	// create a check to see if a user has that specific UUID
	foundUsername, err := client.User.
		Query().
		Where(user.UsernameEQ(newUser.UserName)).Only(ctx)

	if foundUsername != nil {
		log.Println("There is a user with that username")
		return true, err
	}

	// look password
	foundUserPassword, err := client.User.
		Query().
		Where(user.PasswordEQ(newUser.Password)).Only(ctx)

	if foundUserPassword != nil {
		log.Println("There is a user with that password ")
		return true, err
	}

	return false, nil
}

// Updating a user
func UpdateUser(ctx context.Context, c *fiber.Ctx, client *ent.Client) error {
	// create the struct for the user
	newUser := new(M.User)

	if err := c.BodyParser(newUser); err != nil {
		return err
	}

	// update if no issues
	err := client.User.
		Update().
		Where(
			user.UserUUIDEQ(newUser.UUID_Id),
		).
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		SetUsername(newUser.UserName).
		SetPassword(newUser.Password).
		Exec(ctx)
		// check the save status
	if err != nil {
		return fmt.Errorf("Failed to creating a new user: %w", err)
	}
	return nil
}

// Deleting a user

// Grab a group of users

// Grab all users in the database
