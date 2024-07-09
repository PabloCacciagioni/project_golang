package routes

import (
	"strconv"

	"github.com/PabloCacciagioni/project_golang.git/database"
	"github.com/PabloCacciagioni/project_golang.git/models"

	"github.com/gofiber/fiber/v2"
)

func AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
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
