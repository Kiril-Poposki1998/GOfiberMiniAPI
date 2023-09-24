package main

import (
	"log"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup_routes(app)
	log.Fatal(app.Listen(":8080"))
}
