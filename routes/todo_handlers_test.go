package routes_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/PabloCacciagioni/project_golang/models"
	"github.com/PabloCacciagioni/project_golang/routes"
)

var createdTodoID uint64

func SetupApp() *fiber.App {
	app := fiber.New()
	routes.SetupRoutes(app)

	return app
}

func TestAddTodo_Succes(t *testing.T) {
	app := SetupApp()

	title := fmt.Sprintf(`{"title":"My TODO %s"}`, gofakeit.UUID())
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(title))
	req.Header.Set("Content-type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var todo models.Todo
	json.NewDecoder(resp.Body).Decode(&todo)
	createdTodoID = todo.ID
}

func TestListTodos(t *testing.T) {
	app := SetupApp()

	assert.NotZero(t, createdTodoID, "The todo ID should not be zero. Make sure TestAddTodo_Success runs before this test.")

	req := httptest.NewRequest("GET", "/todos", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var todos []models.Todo
	json.NewDecoder(resp.Body).Decode(&todos)

	assert.Greater(t, len(todos), 0)
}

func TestGetTodo_Succes(t *testing.T) {
	app := SetupApp()

	assert.NotZero(t, createdTodoID, "The todo ID should not be zero. Make sure TestAddTodo_Success runs before this test.")

	req := httptest.NewRequest("GET", "/todos/"+strconv.Itoa(int(createdTodoID)), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestUpdateTodo_Succes(t *testing.T) {
	app := SetupApp()

	assert.NotZero(t, createdTodoID, "The todo ID should not be zero. Make sure TestAddTodo_Success runs before this test.")

	req := httptest.NewRequest("PUT", "/todos/"+strconv.Itoa(int(createdTodoID)), strings.NewReader(`{"title":"Updated Todo"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestDeleteTodo_Succes(t *testing.T) {
	app := SetupApp()

	assert.NotZero(t, createdTodoID, "The todo ID should not be zero. Make sure TestAddTodo_Success runs before this test.")

	req := httptest.NewRequest("DELETE", "/todos/"+strconv.Itoa(int(createdTodoID)), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
