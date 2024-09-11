package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
}

func CreateUser(user *User, db *gorm.DB) error {
	err := db.Create(&user).Error
	return err
}

func GetUserById(id uint64, db *gorm.DB) (user *User, err error) {
	user = &User{ID: id}
	err = db.First(user).Error
	return user, err
}

func CheckUserName(username string, db *gorm.DB) (User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
