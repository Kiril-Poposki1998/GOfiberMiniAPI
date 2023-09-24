package main

import (
	"log"
	"os"

	"github.com/gofiber/template/html/v2"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	routes.Setup_routes(app)
	listen_port := ":" + os.Getenv("PORT")
	log.Fatal(app.Listen(listen_port))
}
