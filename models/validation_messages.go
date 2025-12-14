package models


type NewUserValidationError struct{
	SignUpField			string	`json:"field"`
	ValidationMessage	string `json:"string"`
}

