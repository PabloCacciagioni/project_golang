package routes

import (
	"io"
	"net/http"
	"testing"

	"github.com/PabloCacciagioni/project_golang.git/database"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setUpRoute(app *fiber.App) {
	app.Get("/book/:id", GetTodo)
	app.Post("/book", AddTodo)
	app.Put("/book/:id", Update)
	app.Delete("/book/:id", Delete)
}

func Setup() *fiber.App {
	app := fiber.New()
	database.ConnectDb()
	setUpRoute(app)
	return app
}

func TestIndexRoute(t *testing.T) {

	tests := []struct {
		description string

		route string

		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "index route",
			route:         "/",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	app := Setup()

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := io.ReadAll(res.Body)

		assert.Nilf(t, err, test.description)

		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
