package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSaveTodoItem(t *testing.T) {
	dsn := "todouser:todopass@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&Todo{}); err != nil {
		t.Fatalf("Error during auto migration: %v", err)
	}

	todo := Todo{TodoTitle: "Test Todo", Description: "This is a test todo item"}

	if err := db.Create(&todo).Error; err != nil {
		t.Fatalf("Error saving todo item: %v", err)
	}

	var retrievedTodo Todo
	if err := db.First(&retrievedTodo, todo.ID).Error; err != nil {
		t.Fatalf("Error retrieving todo item: %v", err)
	}

	assert.Equal(t, todo.TodoTitle, retrievedTodo.TodoTitle)
	assert.Equal(t, todo.Description, retrievedTodo.Description)
	assert.NotEqual(t, uint(0), retrievedTodo.ID)
}
