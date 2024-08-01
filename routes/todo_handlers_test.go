package routes_test

import (
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/PabloCacciagioni/project_golang/database"
	"github.com/PabloCacciagioni/project_golang/models"
	"github.com/PabloCacciagioni/project_golang/routes"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	db := database.ConnectDb()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.SetupRoutes(app)

	return app
}

func TestListTodos(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("GET", "/todos", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestAddTodo_Succes(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("POST", "/todos", strings.NewReader(`{"title":"Test Todo"}`))
	req.Header.Set("Content-type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestGetTodo_Succes(t *testing.T) {
	app := SetupApp()

	db := database.ConnectDb()
	todo := models.Todo{Title: "Test Todo"}
	db.Create(&todo)

	req := httptest.NewRequest("GET", "/todos/"+strconv.Itoa(int(todo.ID)), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestUpdateTodo_Succes(t *testing.T) {
	app := SetupApp()

	db := database.ConnectDb()
	todo := models.Todo{Title: "Test Todo"}
	db.Create(&todo)

	req := httptest.NewRequest("PUT", "/todos/"+strconv.Itoa(int(todo.ID)), strings.NewReader(`{"title":"Updated Todo"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestDeleteTodo_Succes(t *testing.T) {
	app := SetupApp()

	db := database.ConnectDb()
	todo := models.Todo{Title: "Test Todo"}
	db.Create(&todo)

	req := httptest.NewRequest("DELETE", "/todos/"+strconv.Itoa(int(todo.ID)), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
