package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	Title       string `gorm:"type:VARCHAR(255);not null"`
	Description string `gorm:"type:TEXT"`
}
