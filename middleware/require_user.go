package middleware

import (
	"github.com/gofiber/fiber/v2"
	BLLAuth "teamplayer/bll/auth"
)



func RequireAuth(c *fiber.Ctx) error {
	if BLLAuth.CurrentUser(c) == nil {
		//loop guard
		if c.Path() == "/" { // or "/login"
            return c.Next()
        }
        return c.Redirect("/")
	}
	return c.Next()
}