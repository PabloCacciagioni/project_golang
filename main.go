package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `gorm:"primaryKey;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	Title       string `gorm:"type:VARCHAR(255);not null"`
	Description string `gorm:"type:TEXT"`
}

func main() {
	app := fiber.New()

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Listen(":3000")

	dsn := "todouser:todopass@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&Todo{}); err != nil {
		log.Fatalf("Error during auto migration: %v", err)
	}
}
