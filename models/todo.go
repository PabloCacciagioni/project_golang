package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	ID          uint64 `json:"id" gorm:"primaryKey"`
	CreatedBy   uint64 `json:"created_by"`
	Title       string `json:"title" gorm:"index:idx_todo_title,unique" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=1000"`
}

func (t *Todo) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func GetTodo(id, createdBy uint64, db *gorm.DB) (*Todo, error) {
	todo := &Todo{}

	err := db.Where("id = ? AND created_by = ?", id, createdBy).First(todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return todo, nil
}

func (t *Todo) ListTodo(createdBy uint64, db *gorm.DB) ([]Todo, error) {
	var todos []Todo

	err := db.Where("created_by = ?", t.CreatedBy).Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, err
}

func (t *Todo) Create(db *gorm.DB) error {
	if t.CreatedBy == 0 {
		return gorm.ErrInvalidData
	}

	return db.Create(t).Error
}

func (t *Todo) Update(db *gorm.DB) error {
	existingTodo := &Todo{}
	err := db.Where("id = ? AND created_by = ?", t.ID, t.CreatedBy).First(existingTodo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return db.Save(t).Error
}

func (t *Todo) Delete(db *gorm.DB) error {
	existingTodo := &Todo{}
	err := db.Where("id = ? AND created_by = ?", t.ID, t.CreatedBy).First(existingTodo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return db.Delete(existingTodo).Error
}
