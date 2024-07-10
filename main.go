package main

import (
	"log"

	"github.com/PabloCacciagioni/project_golang.git/config"
	"github.com/PabloCacciagioni/project_golang.git/models"
	"github.com/PabloCacciagioni/project_golang.git/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() (*gorm.DB, error) {
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

func main() {
	app := fiber.New()

	db, err := initDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
