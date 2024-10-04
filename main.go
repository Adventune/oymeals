package main

import (
	"os"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"

	"github.com/adventune/oymeals/handlers"
)

func main() {
	if os.Getenv("STATIC") != "" {
		static()
	} else {
		dynamic()
	}
}

func dynamic() {
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

func static() {
	// Parse the template from the file
	tmpl, err := template.ParseFiles(
		"views/index.tmpl.html",
	)
	if err != nil {
		panic(err)
	}

	// Create or open the file to write the rendered output
	outfile := "meals.html"
	if os.Getenv("OUTFILE") != "" {
		outfile = os.Getenv("OUTFILE")
	}
	f, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	title := "OYMeals"
	if os.Getenv("TITLE") != "" {
		title = os.Getenv("TITLE")
	}

	date := time.Now()
	if date.Hour() >= 17 {
		date = date.Add(24 * time.Hour)
	}

	// Render the template into the file
	err = tmpl.Execute(f, handlers.RenderData{
		Title:       title,
		Restaurants: handlers.Data(),
		Date:        date.Format("02.01.2006"),
	})
	if err != nil {
		panic(err)
	}
}
