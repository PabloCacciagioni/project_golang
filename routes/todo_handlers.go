package routes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PabloCacciagioni/project_golang/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
	"gorm.io/gorm"
)

func ListTodos(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	todo := new(models.Todo)
	todo.CreatedBy = c.Locals("userId").(uint64)

	fm := fiber.Map{
		"type": "alert-error",
	}

	todosSlice, err := todo.ListTodo(todo.CreatedBy, db)
	if err != nil {
		fm["message"] = fmt.Sprintf("something went wrong: %s", err)
		return flash.WithError(c, fm).Redirect("/todo/list")
	}

	return c.Render("todo/index", fiber.Map{
		"Page":     "Task List",
		"Todos":    todosSlice,
		"UserId":   c.Locals("userId").(uint64),
		"Username": c.Locals("username").(string),
		"Message":  flash.Get(c),
	})
}

func AddTodo(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	if c.Method() == "POST" {
		fm := fiber.Map{
			"type": "alert-error",
		}

		newTodo := new(models.Todo)
		newTodo.CreatedBy = c.Locals("userId").(uint64)
		newTodo.Title = strings.TrimSpace(c.FormValue("title"))
		newTodo.Description = strings.TrimSpace(c.FormValue("description"))

		if err := newTodo.Create(db); err != nil {
			fm["message"] = fmt.Sprintf("something went wrong: %s", err)
			return flash.WithError(c, fm).Redirect("/todo/list")
		}

		return c.Redirect("/todo/list")
	}

	return c.Render("todo/edit", fiber.Map{
		"Page":     "Create Todo",
		"UserId":   c.Locals("userId").(uint64),
		"Username": c.Locals("username").(string),
	})
}

func GetTodo(c *fiber.Ctx) error {
	// First we convert the string to the proper unsigned int
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	db := c.Locals("db").(*gorm.DB)

	createdBy := c.Locals("userId").(uint64)

	todo, err := models.GetTodo(id, createdBy, db)
	if err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	db := c.Locals("db").(*gorm.DB)
	userId := c.Locals("userId").(uint64)

	todo, err := models.GetTodo(id, userId, db)
	if err != nil {
		fm := fiber.Map{
			"type":    "alert-error",
			"message": fmt.Sprintf("something went wrong: %s", err),
		}
		return flash.WithError(c, fm).Redirect("/todo/list")
	}

	if c.Method() == "POST" {
		todo.Title = strings.TrimSpace(c.FormValue("title"))
		todo.Description = strings.TrimSpace(c.FormValue("description"))

		if err := todo.Validate(); err != nil {
			fm := fiber.Map{
				"type":    "alert-error",
				"message": fmt.Sprintf("Validation failed: %s", err),
			}
			return flash.WithError(c, fm).Redirect("/todo/edit/%d", int(todo.ID))
		}

		if err := todo.Update(db); err != nil {
			fm := fiber.Map{
				"type":    "alert-error",
				"message": fmt.Sprintf("something went wrong: %s", err),
			}
			return flash.WithError(c, fm).Redirect("todo/list")
		}

		fm := fiber.Map{
			"type":    "alert-succes",
			"message": "Task successfully updated!!",
		}

		return flash.WithSuccess(c, fm).Redirect("/todo/list")
	}

	return c.Render("todo/edit", fiber.Map{
		"Page":     "Edit todo",
		"Todo":     todo,
		"UserId":   userId,
		"Username": c.Locals("username").(string),
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.ErrUnprocessableEntity.Code, err.Error())
	}

	db := c.Locals("db").(*gorm.DB)
	userId := c.Locals("userId").(uint64)

	todo, err := models.GetTodo(id, userId, db)
	if err != nil {
		fm := fiber.Map{
			"type":    "alert-error",
			"message": fmt.Sprintf("something went wrong: %s", err),
		}
		return flash.WithError(c, fm).Redirect("/todo/list", fiber.StatusSeeOther)
	}

	if err := todo.Delete(db); err != nil {
		fm := fiber.Map{
			"type":    "alert-error",
			"message": fmt.Sprintf("something went wrong: %s", err),
		}
		return flash.WithError(c, fm).Redirect("/todo/list", fiber.StatusSeeOther)
	}

	fm := fiber.Map{
		"type":    "alert-success",
		"message": "Task successfully deleted",
	}

	return flash.WithSuccess(c, fm).Redirect("/todo/list", fiber.StatusSeeOther)
}
