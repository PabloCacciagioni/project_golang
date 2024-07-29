package main

import (
	"log"

	"github.com/PabloCacciagioni/project_golang/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8000"))
}
