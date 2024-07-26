package routes

import (
	"strconv"

	"github.com/PabloCacciagioni/project_golang.git/database"
	"github.com/gofiber/fiber/v2"

	"github.com/PabloCacciagioni/project_golang.git/models"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/todo", AddTodo)
	app.Get("/todo/:id", GetTodo)
	app.Put("/todo/:id", Update)
	app.Delete("/todo/:id", Delete)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}

func AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validateStruct(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&todo)
	return c.Status(200).JSON(todo)
}

func GetTodo(c *fiber.Ctx) error {
	todos := []models.Todo{}
	database.DBConn.First(&todos, c.Params("id"))
	return c.Status(200).JSON(todos)
}

func Update(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validateStruct(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	id, _ := strconv.Atoi(c.Params("id"))
	database.DBConn.Model(&models.Todo{}).Where("id = ?", id).Update("title", todo.Title)
	return c.Status(200).JSON("updated")
}

func Delete(c *fiber.Ctx) error {
	todo := new(models.Todo)
	id, _ := strconv.Atoi(c.Params("id"))
	database.DBConn.Where("id = ?", id).Delete(&todo)
	return c.Status(200).JSON("deleted")
}
