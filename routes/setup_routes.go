package routes

import (
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/controller"
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup_routes(app *fiber.App) {
	// API Route
	api := app.Group("/api/")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is up")
	})
	api.Get("/person", controller.GetPersons)
	api.Get("/person/:id", controller.GetPerson)
	api.Delete("/person/:id", controller.DeletePerson)
	api.Put("/person/:id", controller.UpdatePerson)
	api.Post("/person", controller.SetPerson)
	// AUTH Route
	auth := app.Group("/auth/")
	auth.Post("/register", middleware.RegisterUser)
	auth.Post("/login", middleware.LoginUser)
	auth.Get("/logout", middleware.LogoutUser)
}
