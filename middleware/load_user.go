package middleware


import (
	"fmt"
	"teamplayer/ent"
	BLL "teamplayer/bll/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	
	
)


func LoadUserSession(client *ent.Client, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		rawUUID := sess.Get("user_uuid")
		if rawUUID == nil {
			return c.Next()
		}

		uuidStr, ok := rawUUID.(string)
		if !ok || uuidStr == "" {
			sess.Delete("user_uuid")
			return c.Next()
		}

		userUUID, err := uuid.Parse(uuidStr)
		if err != nil {
			sess.Delete("user_uuid")
			return c.Next()
		}

		currentPubProfile, err := BLL.ValidateUserLoad(c.Context(), client, userUUID)
		if err != nil {
			return fmt.Errorf("load user: %w", err)
		}

		if currentPubProfile == nil {
			sess.Delete("user_uuid")
			return c.Next()
		}

		c.Locals("currentUser", currentPubProfile)
		return c.Next()
	}
}