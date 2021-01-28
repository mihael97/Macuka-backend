package services

import (
	"macuka-backend/src/database"
	"macuka-backend/src/models"
	"strconv"
)

func CreateCar(params map[string]string) (*models.Car, error) {
	db := database.GetDatabase()
	miles, err := strconv.ParseInt(params["miles"], 10, 32)
	if err != nil {
		return nil, err
	}
	car := models.Car{
		Miles:             miles,
		RegistrationPlate: params["registration"],
	}
	db.Create(&car)
	return &car, nil
}

func GetCars(idStr string) (interface{}, error) {
	db := database.GetDatabase()
	var cars []models.Car
	if len(idStr) == 0 {
		db.Find(&cars)
		return cars, nil
	} else {
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return nil, err
		}
		db.Find(&cars, id)
		return cars[0], nil
	}
}
