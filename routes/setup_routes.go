package routes

import "github.com/gofiber/fiber/v2"

func Setup_routes(app *fiber.App) fiber.Router {
	setup := app.Group("/api/")
	return setup
}
