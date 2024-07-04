package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `gorm:"primaryKey;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	TodoTitle   string `gorm:"type:VARCHAR(255);not null"`
	Description string `gorm:"type:TEXT"`
}

var db *gorm.DB

func initDB() {
	dsn := "todouser:todopass@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&Todo{})
}

func (t *Todo) Create() error {
	if t.TodoTitle == "" {
		return fmt.Errorf("todo_title cannot be empty")
	}
	return db.Create(t).Error
}

func (t *Todo) Update() error {
	return db.Save(t).Error
}

func (t *Todo) Delete() error {
	return db.Delete(t).Error
}

func (t *Todo) Read(id uint) error {
	return db.First(t, id).Error
}

func ListTodos() ([]Todo, error) {
	var todos []Todo
	result := db.Find(&todos)
	return todos, result.Error
}
func main() {
	app := fiber.New()

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Listen(":3000")

	initDB()

	todo := &Todo{TodoTitle: "Learn Go", Description: "Complete the Go tutorial"}
	err := todo.Create()
	if err != nil {
		log.Fatalf("Error creating todo: %v", err)
	}

	todo.Description = "Complete the Go tutorial and practice"
	err = todo.Update()
	if err != nil {
		log.Fatalf("Error updating todo: %v", err)
	}

	err = todo.Read(todo.ID)
	if err != nil {
		log.Fatalf("Error reading todo: %v", err)
	}

	todos, err := ListTodos()
	if err != nil {
		log.Fatalf("Error listing todos: %v", err)
	}
	for _, t := range todos {
		log.Printf("Todo: %v", t)
	}

	err = todo.Delete()
	if err != nil {
		log.Fatalf("Error deleting todo: %v", err)
	}

}
