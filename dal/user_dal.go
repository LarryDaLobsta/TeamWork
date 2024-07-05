package dal

import (
    "fmt"
    "context"
    "log"
    "teamplayer/ent"
    _ "github.com/lib/pq"
    "github.com/gofiber/fiber/v2"
    M "teamplayer/models"
)


// Creating a user
// Developer note need to check the new addition of a user
// 	Check to see if a user has the same user name and password
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	log.Println("Creating a user")
	newUser, err := client.User.
	Create().
	SetFirstName("Test").
	SetLastName("TestLast").
	SetUsername("test_guy12").
	SetPassword("12345asdfasjdkf4").
	Save(ctx)
	
	// check to make sure the save is successful
	if err != nil {
		log.Println("Failed to create a new user")
		return nil, fmt.Errorf("Failed to creating a new user: %w", err)
	}
	log.Println("User created successfully: ", newUser)
	return newUser, nil
}


func CheckUser(c *fiber.Ctx, client *ent.Client) (bool, error)  {
	// Checks to see if a user already with a password or username provided by a new user

	// create a new user struct
	var newUser = new(M.User)

	// grab and put in struct
	// return false if not able to do it 
	if err := c.BodyParser(newUser); err != nil {
		return false, err
	}





	return true, nil
}




// Finding a single user



// Updating a user



// Deleting a user




// Grab a group of users




// Grab all users in the database

