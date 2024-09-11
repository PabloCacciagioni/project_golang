package models_test

import (
	"testing"

	"github.com/PabloCacciagioni/project_golang/database"
	"github.com/PabloCacciagioni/project_golang/models"
	"github.com/brianvoe/gofakeit/v7"
)

func TestValidation(t *testing.T) {
	todo := models.Todo{
		Title:       "My TODO" + gofakeit.UUID(),
		Description: gofakeit.LoremIpsumSentence(10),
	}

	// Here we check that by default the above thing shold pass
	if todo.Validate() != nil {
		t.Errorf("Expected todo to be valid, but got invalid with error: %v\n", todo.Validate())
	}

	todo = models.Todo{
		Title:       "",
		Description: gofakeit.LoremIpsumSentence(10),
	}

	if todo.Validate() == nil {
		t.Error("Expected todo to be invalid")
	}
}

func TestCreateTodo(t *testing.T) {
	// Ensure db is initialized
	db := database.ConnectDb()

	todo := models.Todo{
		CreatedBy:   1,
		Title:       gofakeit.LoremIpsumWord() + gofakeit.UUID(),
		Description: gofakeit.LoremIpsumSentence(5),
	}

	err := todo.Create(db)
	if err != nil {
		t.Errorf("`todo.Create()` was expected to succeed but failed with err: %v\n", err)
	}

	err = todo.Create(db)
	if err == nil {
		t.Error("`todo.Create()` was expected to fail because of duplicated title")
	}
}

func TestUpdateTodo(t *testing.T) {
	// Ensure db is initialized
	db := database.ConnectDb()

	todo := models.Todo{
		CreatedBy:   1,
		Title:       gofakeit.LoremIpsumWord() + gofakeit.UUID(),
		Description: gofakeit.LoremIpsumSentence(5),
	}

	err := todo.Create(db)
	if err != nil {
		t.Errorf("`todo.Create()` was expected to succeed but failed with err: %v\n", err)
	}

	todo.Title = gofakeit.LoremIpsumWord() + gofakeit.UUID()
	err = todo.Update(db)
	if err != nil {
		t.Errorf("`todo.Update()` was expected to succeed but failed with err: %v\n", err)
	}
}

func TestDeleteTodo(t *testing.T) {
	// Ensure db is initialized
	db := database.ConnectDb()

	todo := models.Todo{
		CreatedBy:   1,
		Title:       gofakeit.LoremIpsumWord() + gofakeit.UUID(),
		Description: gofakeit.LoremIpsumSentence(5),
	}

	err := todo.Create(db)
	if err != nil {
		t.Errorf("`todo.Create()` was expected to succeed but failed with err: %v\n", err)
	}

	err = todo.Delete(db)
	if err != nil {
		t.Errorf("`todo.Delete()` was expected to succeed but failed with err: %v\n", err)
	}
}

func TestGetTodo(t *testing.T) {
	// Ensure db is initialized
	db := database.ConnectDb()

	todo := models.Todo{
		CreatedBy:   1,
		Title:       gofakeit.LoremIpsumWord() + gofakeit.UUID(),
		Description: gofakeit.LoremIpsumSentence(5),
	}

	err := todo.Create(db)
	if err != nil {
		t.Errorf("`todo.Create()` was expected to succeed but failed with err: %v\n", err)
	}

	anotherTodo, err := models.GetTodo(todo.ID, todo.CreatedBy, db)
	if err != nil {
		t.Errorf("`models.GetTodo()` was expected to succeed but failed with err: %v\n", err)
	}

	if todo != *anotherTodo {
		t.Error("both todos were expected to be equal")
	}
}

func TestListTodo(t *testing.T) {
	// Ensure db is initialized
	db := database.ConnectDb()

	todo := models.Todo{
		CreatedBy:   1,
		Title:       gofakeit.LoremIpsumWord() + gofakeit.UUID(),
		Description: gofakeit.LoremIpsumSentence(5),
	}

	err := todo.Create(db)
	if err != nil {
		t.Errorf("`todo.Create()` was expected to succeed but failed with err: %v\n", err)
	}

	allTodos, err := todo.ListTodo(todo.CreatedBy, db)
	if err != nil {
		t.Errorf("`models.ListTodos()` was expected to succeed but failed with err: %v\n", err)
	}

	if len(allTodos) == 0 {
		t.Error("there should be at least one todo that we just created")
	}
}
