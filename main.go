package main

import (
	"log"
	"os"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/database"
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.Connect()
	database.ConnectRedis()
	app := fiber.New(fiber.Config{
		Views: html.New("/src/views", ".html"),
	})
	routes.Setup_routes(app)
	listen_port := ":" + os.Getenv("APP_PORT")
	log.Fatal(app.Listen(listen_port))
}
