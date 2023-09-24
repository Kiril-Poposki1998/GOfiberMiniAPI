package main

import (
	"log"
	"os"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup_routes(app)
	listen_port := ":" + os.Getenv("PORT")
	log.Fatal(app.Listen(listen_port))
}
