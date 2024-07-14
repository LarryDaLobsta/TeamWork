package dal

import (
    "fmt"
    "context"
    "log"
    "teamplayer/ent"
    "teamplayer/ent/user"
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
	SaveX(ctx) // See if different type of save occurs do to panic 
	
	// check to make sure the save is successful
	if err != nil {
		log.Println("Failed to create a new user")
		return nil, fmt.Errorf("Failed to creating a new user: %w", err)
	}
	log.Println("User created successfully: ", newUser)
	return newUser, nil
}


// returns true if user found or false if user not found
func CheckUser(ctx context.Context, c *fiber.Ctx, client *ent.Client) (bool, error)  {
	// Checks to see if a user already with a password or username provided by a new user

	// create a new user struct
	var newUser = new(M.User)

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

	//look password
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
func UpdateUser(ctx context.Context, c *fiber.Ctx, client *ent.Client) (bool, error){

	// create the struct for the user
	var newUser = new(M.User)

	if err := c.BodyParser(newUser); err != nil {
		return false, err
	}

	// update if no issues
	err := client.User.
	UpdateOneID(M.UserUUID).
	SetFirstName(M.User.FirstName).
	SetLastName(M.User.LastName).
	SetUserName(M.User.Username).
	SetPassword(M.User.Password).
	SaveX(ctx)

	// check the save status
	if err != nil {
		log.Println("Failed to update the user")
		return false, fmt.Errorf("Failed to creating a new user: %w", err)
	}

	// may want to create a save status function for the crud functionality
	log.Println("User updated successfully")
	return true, err

}



// Deleting a user




// Grab a group of users




// Grab all users in the database

