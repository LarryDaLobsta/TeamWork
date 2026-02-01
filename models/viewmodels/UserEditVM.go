package models

import (
	models "teamplayer/models"
)

type EditUserViewModel struct {
	Title          string
	Form           models.PublicUserProfile
	UserErrors     models.EditUserErrors
	SystemError    models.ServerErrors
	SuccessMessage string
}
