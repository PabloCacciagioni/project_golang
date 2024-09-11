package routes_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"

	"github.com/PabloCacciagioni/project_golang/routes"
)

func TestHandleViewHome(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status code %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestHandleViewLogin(t *testing.T) {
	// Create a new instance of Fiber app
	app := fiber.New()

	routes.SetupRoutes(app)

	formData := url.Values{}
	formData.Set("username", "testuser")

	req := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := app.Test(req)

	// Check that there were no errors during the request
	require.NoError(t, err)

	// Check for redirection
	require.Equal(t, fiber.StatusFound, resp.StatusCode)
	location := resp.Header.Get(fiber.HeaderLocation)
	require.Equal(t, "/todo/list", location)
}
