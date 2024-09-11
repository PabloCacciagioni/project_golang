package database

import (
	"log"
	"os"

	"github.com/PabloCacciagioni/project_golang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultDSN = "todouser:todopass@/tododb?charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify&autocommit=true&collation=utf8mb4_unicode_ci"

func ConnectDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(generateDSNFromEnv()), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to the database with error: " + err.Error())
	}

	// TODO: Fix this crap
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Todo{}, &models.User{})

	return db
}

func generateDSNFromEnv() string {
	mysqlURI := os.Getenv("MYSQL_URI")
	if mysqlURI == "" {
		mysqlURI = defaultDSN
	}
	return mysqlURI
}
