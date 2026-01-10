package models

import (
	_ "fmt"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

type UserRecord struct {
	// adding to the ent struct for dealing with application post requests
	UUID_Id      uuid.UUID 
	FirstName    string    
	LastName     string    
	UserName     string
	Email		 string    
	PasswordHash string    
}

type PublicUserProfile struct {
	UUID_Id      uuid.UUID	
	FirstName    string		  
	LastName     string    	
	UserName     string		
	Email		 string    	
}

