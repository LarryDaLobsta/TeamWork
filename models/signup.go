package models


type UserSignUp struct {
    FirstName string	`json:"first_name"`
    LastName  string	`json:"last_name"`
    Username  string	`json:"username"`
    Email     string	`json:"email"`
    Password  string	`json:"password"`
}


type UserProfile struct {
	id		  int		`json:"id"`
    FirstName string	`json:"first_name"`
    LastName  string	`json:"last_name"`
    Username  string	`json:"username"`	
    Email     string	`json:"email"`
}