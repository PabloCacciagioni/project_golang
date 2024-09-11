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

var (
	title         = gofakeit.Sentence(3)
	description   = gofakeit.Paragraph(1, 5, 10, ".")
	createdTodoID uint64
)

func SetupApp() *fiber.App {
	app := fiber.New()
	routes.SetupRoutes(app)

	return app
}

func TestAddTodo_Success(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	body := fmt.Sprintf("title=%s&description=%s", title, description)
	req := httptest.NewRequest("POST", "/todo/create", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusFound, resp.StatusCode)
}

func TestListTodos(t *testing.T) {
	app := SetupApp()

	body := fmt.Sprintf("title=%s&description=%s", title, description)
	createReq := httptest.NewRequest("POST", "/todo/create", strings.NewReader(body))
	createReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	createResp, err := app.Test(createReq)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusFound, createResp.StatusCode)

	listReq := httptest.NewRequest("GET", "/todo/list", nil)

	listResp, err := app.Test(listReq)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, listResp.StatusCode)
}

func TestGetTodo_Succes(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("GET", "/todos/"+strconv.Itoa(int(createdTodoID)), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusFound, resp.StatusCode)
}

func TestUpdateTodo_Succes(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest("PUT", "/todo/edit"+strconv.Itoa(int(createdTodoID)), strings.NewReader(`{"title":"Updated Todo"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusFound, resp.StatusCode)

	var updatedTodo models.Todo
	json.NewDecoder(resp.Body).Decode(&updatedTodo)

	assert.Equal(t, "Updated Todo", updatedTodo.Title)
	assert.Equal(t, createdTodoID, updatedTodo.ID)
}

func TestDeleteTodo_Succes(t *testing.T) {
	app := SetupApp()

	newTodoTitle := fmt.Sprintf("My TODO %s", gofakeit.UUID())
	createReq := httptest.NewRequest("POST", "/todo/create", strings.NewReader(fmt.Sprintf(`{"title":"%s"}`, newTodoTitle)))
	createReq.Header.Set("Content-Type", "application/json")
	createResp, _ := app.Test(createReq)
	assert.Equal(t, fiber.StatusFound, createResp.StatusCode)

	var newTodo models.Todo
	json.NewDecoder(createResp.Body).Decode(&newTodo)

	deleteReq := httptest.NewRequest("DELETE", "/todo/delete/"+strconv.Itoa(int(newTodo.ID)), nil)
	deleteResp, _ := app.Test(deleteReq)
	assert.Equal(t, fiber.StatusFound, deleteResp.StatusCode)

	getReq := httptest.NewRequest("GET", "/todos/"+strconv.Itoa(int(newTodo.ID)), nil)
	getResp, err := app.Test(getReq)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusFound, getResp.StatusCode)
}
