package routes

import (
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup_routes(app *fiber.App) {
	api := app.Group("/api/")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is up")
	})
	api.Get("/person", controller.GetPerson)
	api.Post("/person", controller.SetPerson)
}
