package services

import (
	"macuka-backend/src/database"
	"macuka-backend/src/models"
	"strconv"
)

func DeleteCar(id string) error {
	idValue, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	db := database.GetDatabase()
	db.Exec("DELETE FROM trips WHERE car=?", uint(idValue))
	db.Exec("DELETE FROM cars WHERE id=?", uint(idValue))
	return nil
}

func CreateCar(params map[string]string) (*models.Car, error) {
	db := database.GetDatabase()
	miles, err := strconv.ParseUint(params["miles"], 10, 32)
	if err != nil {
		return nil, err
	}
	year, err := strconv.ParseUint(params["productionYear"], 10, 32)
	if err != nil {
		return nil, err
	}
	car := models.Car{
		Name:              params["name"],
		Miles:             uint(miles),
		RegistrationPlate: params["registrationPlate"],
		Year:              uint(year),
	}
	db.Create(&car)
	return &car, nil
}

func GetCars(idStr string) (interface{}, error) {
	db := database.GetDatabase()
	var cars []models.Car
	if len(idStr) == 0 {
		db.Raw("SELECT * FROM cars").Find(&cars)
		returnCars := make([]models.Car, 0)
		for _, car := range cars {
			db.Raw("SELECT * FROM trips WHERE car=?", car.Id).Find(&car.Trips)
			returnCars = append(returnCars, car)
		}

		return returnCars, nil
	} else {
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return nil, err
		}
		db.Find(&cars, id)
		db.Raw("SELECT * FROM trips WHERE car=?", cars[0].Id).Find(&cars[0].Trips)
		return cars[0], nil
	}
}
