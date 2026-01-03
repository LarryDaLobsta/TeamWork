package models

import (
	models "teamplayer/models"
)


type SignUpViewModel struct {
	Form 	models.UserSignUp
	Errors 	models.SignUpErrors
}