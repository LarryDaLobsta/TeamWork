package bll

import (
	// "encoding/json"

	"context"
	"teamplayer/ent"
	viewmodels "teamplayer/models/viewmodels"
	// "github.com/gofiber/contrib/websocket"
)

func EditUser(ctx context.Context, client *ent.Client, EditUserVM viewmodels.EditUserViewModel) viewmodels.EditUserViewModel {

	// check to make sure there is no empty view model

	// check to see if the fields are different from the original entries
	// for the fields that are different validate accordingly

	// if email or username fall in that category then use IsAlready exist functions in DAL
	// then validate the if need be

	// return new edit page changes and any errors if any

	return EditUserVM

}
