package dal

import (
    "fmt"
    "context"
    "log"
    "teamplayer/ent"
    _ "github.com/lib/pq"
)


// Creating a user
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
		return nil, fmt.Errorf("Failed to creating a new user: %w", err)
	}
	log.Println("User created successfully: ", newUser)
	return newUser, nil
}







// Finding a single user



// Updating a user



// Deleting a user




// Grab a group of users




// Grab all users in the database

