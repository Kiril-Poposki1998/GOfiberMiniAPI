package main

import (
	"goapp/controller"
	"goapp/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router := routes.Setup_routes(app)
	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})
	router.Get("/person", controller.GetPerson)
	log.Fatal(app.Listen(":8080"))
}
