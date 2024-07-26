package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	ID          uint64 `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"index:idx_todo_title,unique" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=1000"`
}

func (t *Todo) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

// GetTodo returns the requested Todo or error
func GetTodo(id uint64, db *gorm.DB) (todo *Todo, err error) {
	todo = &Todo{ID: id}
	err = db.First(todo).Error
	return todo, err
}

func ListTodos(db *gorm.DB) (todos []Todo, err error) {
	err = db.Find(&todos).Error
	return todos, err
}

func (t *Todo) Create(db *gorm.DB) (err error) {
	return db.Create(t).Error
}

func (t *Todo) Update(db *gorm.DB) (err error) {
	return db.Save(t).Error
}

func (t *Todo) Delete(db *gorm.DB) (err error) {
	return db.Delete(t).Error
}
