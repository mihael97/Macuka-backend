package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"macuka-backend/src/models"
	"os"
)

var database *gorm.DB = nil

func InitializeDatabase() {
	var host = os.Getenv("DB_HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	databasePort := os.Getenv("DB_PORT")
	if len(databasePort) == 0 {
		databasePort = "5432"
	}
	var databaseName = os.Getenv("DB_NAME")
	if len(databaseName) == 0 {
		databaseName = "database"
	}
	var databaseUser = os.Getenv("DB_USER")
	if len(databaseUser) == 0 {
		databaseUser = "user"
	}
	var databasePassword = os.Getenv("DB_PASS")
	if len(databasePassword) == 0 {
		databasePassword = "password"
	}
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, databasePort, databaseUser, databasePassword, databaseName)
	var err error
	database, err = gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	migrateTables()
}

func migrateTables() {
	err := database.AutoMigrate(&models.Car{}, &models.Trip{}, &models.Customer{})
	if err != nil {
		log.Panic(err)
	}
}

func GetDatabase() *gorm.DB {
	if database == nil {
		InitializeDatabase()
	}
	return database
}
