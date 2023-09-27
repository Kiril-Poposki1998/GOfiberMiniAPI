package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/database"
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AuthenticateUser(c *fiber.Ctx) {
	user_id := uuid.New()
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    user_id.String(),
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})
	database.Redis.Set(user_id.String(), []byte("true"), 24*time.Hour)
}

func CheckUser(c *fiber.Ctx) bool {
	auth_cookie := c.Cookies("session_id")
	check, err := database.Redis.Get(auth_cookie)
	if string(check) == "true" {
		return true
	} else if err != nil {
		fmt.Print(err)
	}
	return false
}

func RegisterUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Print(err.Error())
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
	database.Database.Create(user)
	return c.Status(fiber.StatusOK).SendString("User created")
}

func LoginUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		fmt.Print(err.Error())
		return err
	}
	fetch_user := new(model.User)
	database.Database.First(&fetch_user, "Username = ?", user.Username)
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
	if fetch_user.Password == user.Password {
		AuthenticateUser(c)
		return c.Status(fiber.StatusOK).SendString("You have been logged in. You can view each individual user")
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Username or password is not correct")
}

func LogoutUser(c *fiber.Ctx) error {
	auth_cookie := c.Cookies("session_id")
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
	})
	err := database.Redis.Delete(auth_cookie)
	if err != nil {
		panic("Error while deleting session in redis")
	}
	return c.Status(fiber.StatusOK).SendString("You have been logged out")
}
