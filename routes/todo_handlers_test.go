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

	req := httptest.NewRequest("GET", "/todos", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var todos []models.Todo
	json.NewDecoder(resp.Body).Decode(&todos)

	assert.Greater(t, len(todos), 0)

	found := false
	for _, todo := range todos {
		if todo.ID == createdTodoID {
			found = true
			break
		}
	}
	assert.True(t, found, "The todo with the createdTodoID should be listed")
}

func TestGetTodo_Succes(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("GET", "/todos/"+strconv.Itoa(int(createdTodoID)), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestUpdateTodo_Succes(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("PUT", "/todos/"+strconv.Itoa(int(createdTodoID)), strings.NewReader(`{"title":"Updated Todo"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var updatedTodo models.Todo
	json.NewDecoder(resp.Body).Decode(&updatedTodo)

	assert.Equal(t, "Updated Todo", updatedTodo.Title)
	assert.Equal(t, createdTodoID, updatedTodo.ID)
}

func TestDeleteTodo_Succes(t *testing.T) {
	app := SetupApp()

	newTodoTitle := fmt.Sprintf("My TODO %s", gofakeit.UUID())
	createReq := httptest.NewRequest("POST", "/todos", strings.NewReader(fmt.Sprintf(`{"title":"%s"}`, newTodoTitle)))
	createReq.Header.Set("Content-Type", "application/json")
	createResp, _ := app.Test(createReq)
	assert.Equal(t, fiber.StatusCreated, createResp.StatusCode)

	var newTodo models.Todo
	json.NewDecoder(createResp.Body).Decode(&newTodo)

	deleteReq := httptest.NewRequest("DELETE", "/todos/"+strconv.Itoa(int(newTodo.ID)), nil)
	deleteResp, _ := app.Test(deleteReq)
	assert.Equal(t, fiber.StatusOK, deleteResp.StatusCode)

	getReq := httptest.NewRequest("GET", "/todos/"+strconv.Itoa(int(newTodo.ID)), nil)
	getResp, _ := app.Test(getReq)
	assert.Equal(t, fiber.StatusNotFound, getResp.StatusCode)
}
