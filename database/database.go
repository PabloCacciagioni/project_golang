package database

import (
	"log"
	"os"

	"github.com/PabloCacciagioni/project_golang.git/config"
	"github.com/PabloCacciagioni/project_golang.git/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() *gorm.DB {
	dsn := config.GetDBConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Connection error")
		os.Exit(2)
	}

	log.Println("Connected")
	db.AutoMigrate(&models.Todo{})
	DBConn = db
	return DBConn
}
