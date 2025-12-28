package models


type NewUserValidationError struct {
	SignUpField			string	`json:"field"`
	ValidationMessage	string `json:"validation_message"`
}

func ( response *NewUserValidationError) Error() string {
	return response.ValidationMessage
}