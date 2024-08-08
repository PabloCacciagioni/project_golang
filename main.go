package main

import (
	"log"

	"github.com/PabloCacciagioni/project_golang/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("Serving index.html")
		return c.SendFile("./index.html")
	})

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8000"))
}
