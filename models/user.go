package models


import (
   _"fmt"
   _"github.com/google/uuid"
)



type User struct{
	// adding to the ent struct for dealing with application post requests
	FirstName string `json:"firstname"`  
	LastName  string `json:"lastname"`  
	UserName  string `json:"username"`  
	Password  string `json:"password"`  
}

