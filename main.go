package main

import (
	"log"

	"github.com/PabloCacciagioni/project_golang.git/database"
	"github.com/PabloCacciagioni/project_golang.git/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/book/:id", routes.GetTodo)
	app.Post("/book", routes.AddTodo)
	app.Put("/book/:id", routes.Update)
	app.Delete("/book/:id", routes.Delete)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8000"))
}
