package services

import (
	"macuka-backend/src/database"
	"macuka-backend/src/models"
)

func CreateCities(cities []models.City) {
	db := database.GetDatabase()
	db.Create(&cities)
}

func GetCities() []models.City {
	db := database.GetDatabase()
	var cities []models.City
	db.Find(&cities)
	return cities
}
