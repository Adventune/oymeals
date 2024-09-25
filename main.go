package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"

	"mealfetch/handlers"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 4 * 1024 * 1024 * 1024,
	})
	app.Use(cors.New())

	app.Get("/", handlers.Index)
	app.Static("/", "./static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
