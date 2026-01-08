package models

import (
	models "teamplayer/models"
)


type SignUpViewModel struct {
	Title string
	Form 	models.UserSignUp
	Errors 	models.SignUpErrors
}