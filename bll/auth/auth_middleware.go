package bll

import (
	"context"
	"fmt"
	"teamplayer/ent"
	dal "teamplayer/dal"
	models "teamplayer/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	
)



func ValidateUserLoad(ctx context.Context, client *ent.Client, currUserUUID uuid.UUID) (*models.PublicUserProfile, error) {
	
	var currentPubProfile models.PublicUserProfile
	
	if !IsValidUUID(currUserUUID) {
		return nil, nil
	}

	// go to database to grab user
	currentUserRecord, err := dal.GetUserByUUID(ctx, client, currUserUUID)

	if ent.IsNotFound(err) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get user by uuid %w", err)
	}

	// slap the new user into the public profile
	currentPubProfile.UUID_Id =   currentUserRecord.UserUUID
	currentPubProfile.FirstName = currentUserRecord.FirstName
	currentPubProfile.LastName =  currentUserRecord.LastName
	currentPubProfile.UserName =  currentUserRecord.Username
	currentPubProfile.Email =     currentUserRecord.Email

	return &currentPubProfile, nil

}


func CurrentUser(c *fiber.Ctx) *models.PublicUserProfile {
	if u := c.Locals("currentUser"); u != nil {
		return u.(*models.PublicUserProfile)
	}
	return nil
}