package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

type PageData struct {
	Count int
	Title string
	Users []User
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
}

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false

	pageData := PageData{
		Count: 0,
		Title: "Simple counter",
	}

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		pageData.Count++
		// Render index
		return c.Render("index", fiber.Map{
			"Title": pageData.Title,
			"Count": pageData.Count,
		}, "layouts/main")
	})

	// app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Post("/countup", func(c *fiber.Ctx) error {
		pageData.Count++
		return c.Render("partials/counter", fiber.Map{
			"Count": pageData.Count,
		})
	})
	app.Post("/countdown", func(c *fiber.Ctx) error {
		pageData.Count--
		return c.Render("partials/counter", fiber.Map{
			"Count": pageData.Count,
		})
	})
	app.Post("/countreset", func(c *fiber.Ctx) error {
		pageData.Count = 0
		return c.Render("partials/counter", fiber.Map{
			"Count": pageData.Count,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

func newUser(name string, email string) User {
	return User{
		Name:  name,
		Email: email,
	}
}
