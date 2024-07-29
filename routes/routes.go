package routes

import (
	"github.com/PabloCacciagioni/project_golang/database"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := database.ConnectDb()

	// Here we setup our database in the context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Then we just get the routes
	app.Get("/status", GetStatus)
	app.Get("/todos", ListTodos)
	app.Post("/todos", AddTodo)
	app.Get("/todos/:id", GetTodo)
	app.Put("/todos/:id", UpdateTodo)
	app.Delete("/todos/:id", DeleteTodo)
}
