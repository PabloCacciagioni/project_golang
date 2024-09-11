package main

import (
	"log"

	"github.com/PabloCacciagioni/project_golang/database"
	"github.com/PabloCacciagioni/project_golang/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func init() {
	database.ConnectDb()
}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base.layout",
	})

	app.Static("/", "./assets")

	app.Use(logger.New())

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8000"))
}
