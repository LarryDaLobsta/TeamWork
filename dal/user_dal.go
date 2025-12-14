package dal

import (
	"context"
	"fmt"
	"teamplayer/ent"
	M "teamplayer/models"
	_ "github.com/lib/pq"
)

// Creating a user
// Developer note need to check the new addition of a user
//
//	Check to see if a user has the same user name and password
func CreateUser(newUser M.UserRecord, client *ent.Client, ctx context.Context) error {
	// THis should be put in the route that is handling a new user being passed from the client
	// newUser := new(M.User)
	// // grab and put in struct
	// // return false if not able to do it
	// if err := c.BodyParser(newUser); err != nil {
	// 	return err
	// }
	// also need to update all parameters getting issue about 

	// end of that piece

	createdNewUser, err := client.User.Create().
		SetFirstName(newUser.FirstName).
		SetLastName(newUser.LastName).
		SetEmail(newUser.Email).
		SetUsername(newUser.UserName).
		SetPasswordHash(newUser.PasswordHash).
		Save(ctx) // See if different type of save occurs do to panic
	if err != nil {
		return fmt.Errorf("Error: %v %v", err, createdNewUser)
	}
	return nil
}

// returns true if user found or false if user not found
// func CheckUser(ctx context.Context, c *fiber.Ctx, client *ent.Client) error {
// 	// Checks to see if a user already with a password or username provided by a new user

// 	// create a new user struct
// 	newUser := new(M.User)

// 	// grab and put in struct
// 	// return false if not able to do it
// 	if err := c.BodyParser(newUser); err != nil {
// 		return err
// 	}

// 	// check to see if the username and password is in the database

// 	// create a check to see if a user has that specific UUID
// 	foundUsername, err := client.User.
// 		Query().
// 		Where(user.UsernameEQ(newUser.UserName)).Only(ctx)

// 	if foundUsername != nil {
// 		return err
// 	}

// 	// look password
// 	foundUserPassword, err := client.User.
// 		Query().
// 		Where(user.PasswordEQ(newUser.Password)).Only(ctx)

// 	if foundUserPassword != nil {
// 		return err
// 	}

// 	return nil
// }

// // Updating a user
// func UpdateUser(ctx context.Context, c *fiber.Ctx, client *ent.Client) error {
// 	// create the struct for the user
// 	newUser := new(M.User)

// 	if err := c.BodyParser(newUser); err != nil {
// 		return err
// 	}

// 	// update if no issues
// 	err := client.User.
// 		Update().
// 		Where(
// 			user.UserUUIDEQ(newUser.UUID_Id),
// 		).
// 		SetFirstName(newUser.FirstName).
// 		SetLastName(newUser.LastName).
// 		SetUsername(newUser.UserName).
// 		SetPassword(newUser.Password).
// 		Exec(ctx)
// 		// check the save status
// 	if err != nil {
// 		return fmt.Errorf("Failed to creating a new user: %w", err)
// 	}
// 	return nil
// }

// Deleting a user

// Grab a group of users

// Grab all users in the database
