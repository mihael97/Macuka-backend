package services

import (
	"errors"
	"fmt"
	"macuka-backend/src/database"
	"macuka-backend/src/models"
	"strconv"
)

func CreateTrip(tripDto models.TripDto) (*models.Trip, error) {
	start, err := strconv.ParseUint(tripDto.Start, 10, 32)
	if err != nil {
		return nil, nil
	}
	end, err := strconv.ParseUint(tripDto.End, 10, 32)
	if err != nil {
		return nil, nil
	}
	carId, err := strconv.ParseUint(tripDto.Car, 10, 32)
	db := database.GetDatabase()
	var cars []models.Car
	db.Find(&cars, uint(carId))
	if len(cars) == 0 {
		return nil, errors.New(fmt.Sprintf("%s car doesn't exist", tripDto.Car))
	}
	trip := models.Trip{
		Path:  tripDto.ConnectPath(),
		Start: uint(start),
		End:   uint(end),
		Car:   uint(carId),
	}
	db.Create(&trip)
	err = db.Model(&cars[0]).Association("Trips").Append([]models.Trip{trip})
	if err != nil {
		return nil, err
	}
	return &trip, nil
}
