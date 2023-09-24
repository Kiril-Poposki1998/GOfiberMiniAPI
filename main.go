package main

import (
	"goapp/controller"
	"goapp/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := routes.Setup_routes(app)
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})
	api.Get("/person", controller.GetPerson)
	api.Post("/person", controller.SetPerson)
	log.Fatal(app.Listen(":8080"))
}
