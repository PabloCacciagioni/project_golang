package models_test

import (
	"testing"

	"github.com/PabloCacciagioni/project_golang/database"
	"github.com/PabloCacciagioni/project_golang/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db := database.ConnectDb()

	user := &models.User{
		Username: gofakeit.Username(),
	}

	err := models.CreateUser(user, db)
	assert.NoError(t, err, "`CreateUser()` was espected to succed but failed")

	duplicateUser := &models.User{
		Username: user.Username,
	}

	err = models.CreateUser(duplicateUser, db)
	assert.Error(t, err, "`CreateUser()` was expected to fail because of duplicated username")
}

func TestGetUserById(t *testing.T) {
	db := database.ConnectDb()

	user := &models.User{
		Username: gofakeit.Username(),
	}

	err := models.CreateUser(user, db)
	assert.NoError(t, err, "`CreateUser()` was expected to succeed but failed")

	retrievedUser, err := models.GetUserById(user.ID, db)
	assert.NoError(t, err, "`GetUserById()` was expected to succeed but failed")
	assert.Equal(t, user.ID, retrievedUser.ID, "Expected IDs to match")
	assert.Equal(t, user.Username, retrievedUser.Username, "Expected usernames to match")

	nonExistentID := user.ID + 1
	_, err = models.GetUserById(nonExistentID, db)
	assert.Error(t, err, "`GetUserById()` was expected to fail for non-existent user")
	assert.Equal(t, gorm.ErrRecordNotFound, err, "Expected error to be `record not found`")
}

func TestCheckUserName(t *testing.T) {
	db := database.ConnectDb()

	user := &models.User{
		Username: gofakeit.Username(),
	}

	err := models.CreateUser(user, db)
	assert.NoError(t, err, "`CreateUser()` was expected to succeed but failed")

	retrievedUser, err := models.CheckUserName(user.Username, db)
	assert.NoError(t, err, "`CheckUserName()` was expected to succeed but failed")
	assert.Equal(t, user.ID, retrievedUser.ID, "Expected IDs to match")
	assert.Equal(t, user.Username, retrievedUser.Username, "Expected usernames to match")

	nonExistentUsername := gofakeit.Username()
	_, err = models.CheckUserName(nonExistentUsername, db)
	assert.Error(t, err, "`CheckUserName()` was expected to fail for non-existent user")
	assert.Equal(t, gorm.ErrRecordNotFound, err, "Expected error to be `record not found`")
}
