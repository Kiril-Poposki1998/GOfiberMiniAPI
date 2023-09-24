package main

import (
	"goapp/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup_routes(app)
	log.Fatal(app.Listen(":8080"))
}
