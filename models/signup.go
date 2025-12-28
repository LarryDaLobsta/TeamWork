package models


type UserSignUp struct {
    FirstName string	`form:"first_name"`
    LastName  string	`form:"last_name"`
    UserName  string	`form:"username"`
    Email     string	`form:"email"`
    Password  string	`form:"password"`
}


type UserProfile struct {
	id		  int		`form:"id"`
    FirstName string	`form:"first_name"`
    LastName  string	`form:"last_name"`
    UserName  string	`form:"username"`	
    Email     string	`form:"email"`
}