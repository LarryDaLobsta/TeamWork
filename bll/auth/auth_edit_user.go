package bll


import (
	// "encoding/json"
	"strings"
	"strconv"
	"net/mail"
	"unicode"
	"golang.org/x/crypto/bcrypt"
	"errors"
	models "teamplayer/models"
	viewmodels "teamplayer/models/viewmodels"
	"github.com/google/uuid"
	// "github.com/gofiber/contrib/websocket"
	dal "teamplayer/dal"
)




func EditUser(ctx context.Context, client *ent.Clinet, EditUserVM editUserVM) EditUserVM, error {
	
	// check to make sure there is no empty view model

	// check to see if the fields are different from the original entries
		// for the fields that are different validate accordingly

		// if email or username fall in that category then use IsAlready exist functions in DAL
		// then validate the if need be

		// return new edit page changes and any errors if any 


	


}