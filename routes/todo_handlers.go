package routes

import (
	"strconv"

	"github.com/PabloCacciagioni/project_golang/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ListTodos(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	allTodos, err := models.ListTodos(db)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(allTodos)
}

func AddTodo(c *fiber.Ctx) error {
	todo := models.Todo{}

	// First we validate they are sending JSON
	if err := c.BodyParser(&todo); err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	// Then we validate they send the JSON as we expect it
	if err := todo.Validate(); err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	// Then we extract the database from context
	db := c.Locals("db").(*gorm.DB)

	// Then we try to save it and if it fails we send back the message
	if err := todo.Create(db); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	// If everything succeeded we then send the all is fine message
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func GetTodo(c *fiber.Ctx) error {
	// First we convert the string to the proper unsigned int
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	db := c.Locals("db").(*gorm.DB)

	todo, err := models.GetTodo(id, db)
	if err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	// First we convert the string to the proper unsigned int
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}
	db := c.Locals("db").(*gorm.DB)

	todo, err := models.GetTodo(id, db)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	// then we validate they are sending JSON
	if err := c.BodyParser(&todo); err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	// Then we validate they send the JSON as we expect it
	if err := todo.Validate(); err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	// Here we retrieve the current value
	if err := todo.Update(db); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}
	todo := models.Todo{ID: id}

	db := c.Locals("db").(*gorm.DB)

	// Here we delete the corresponding todo
	if err := todo.Delete(db); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"deleted": true,
		"id":      id,
	})
}
