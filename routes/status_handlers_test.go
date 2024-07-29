package routes_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/PabloCacciagioni/project_golang/routes"
	"github.com/gofiber/fiber/v2"
)

func TestGetStatus(t *testing.T) {
	app := fiber.New()

	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/status", nil)

	// Perform the request plain with the app,
	// the second argument is a request latency
	// (set to -1 for no latency)
	resp, _ := app.Test(req, -1)

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status code %v but received %v\n", fiber.StatusOK, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "{\"status\":\"OK\"}" {
		t.Errorf("expected {\"status\":\"OK\"} but received %s", body)
	}
}
