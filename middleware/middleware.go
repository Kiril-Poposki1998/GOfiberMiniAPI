package middleware

import (
	"fmt"
	"time"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AuthenticateUser(c *fiber.Ctx) error {
	user_id := uuid.New()
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    user_id.String(),
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})
	database.Redis.Set(user_id.String(), []byte("true"), 24*time.Hour)
	return c.Status(fiber.StatusOK).SendString("You have been logged in")
}

func CheckUser(c *fiber.Ctx) bool {
	auth_cookie := c.Cookies("session_id")
	check, err := database.Redis.Get(auth_cookie)
	if string(check) == "true" {
		return true
	}
	fmt.Print(err)
	return false
}
