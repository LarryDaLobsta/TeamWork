package models

type SignUpErrors struct {
	FirstName string
	LastName  string
	Email     string
	UserName  string
	Password  string
}

type EditUserErrors struct {
	FirstName string
	LastName  string
	Email     string
	UserName  string
}

type ServerErrors struct {
	SystemError string
}
