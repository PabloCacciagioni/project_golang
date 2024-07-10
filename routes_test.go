package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/PabloCacciagioni/project_golang.git/config"
	"github.com/PabloCacciagioni/project_golang.git/models"
	"github.com/PabloCacciagioni/project_golang.git/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initTestDatabase() (*gorm.DB, error) {
	dsn := config.GetDBConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		return nil, err
	}

	return db, nil
}
func TestIndexRoute(t *testing.T) {
	db, err := initTestDatabase()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	tests := []struct {
		description   string
		route         string
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

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.SetupRoutes(app)

	for _, test := range tests {
		req, _ := http.NewRequest("GET", test.route, nil)
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
