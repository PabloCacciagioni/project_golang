package config

import "fmt"

const (
	DBPort     = "3306"
	DBUser     = "todouser"
	DBPassword = "todopass"
	DBName     = "tododb"
)

func GetDBConnection() string {
	connection := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBPort,
		DBName)
	return connection
}
