package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/PabloCacciagioni/project_golang.git/database"
	"github.com/PabloCacciagioni/project_golang.git/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func AddTodoValidation(todo *models.Todo) error {
	return validate.Struct(todo)
}

func AddTodoHandler(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := AddTodoValidation(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DBConn.Create(&todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

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
	db := c.Locals("db").(*gorm.DB)
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := db.Create(&todo).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(todo)
}

func GetTodo(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	todo := new(models.Todo)
	if err := db.First(&todo, c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(todo)
}

func Update(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	id, _ := strconv.Atoi(c.Params("id"))

	if err := db.Model(&models.Todo{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON("updated")
}

func Delete(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := db.Delete(&models.Todo{}, id).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON("deleted")
}
