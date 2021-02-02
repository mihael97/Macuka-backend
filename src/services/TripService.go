package services

import (
	"errors"
	"fmt"
	"macuka-backend/src/database"
	"macuka-backend/src/models"
	"strconv"
	"time"
)

const (
	DateFormat = "2006-01-02"
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
	date, err := time.Parse(DateFormat, tripDto.Date)
	if err != nil {
		return nil, err
	}
	trip := models.Trip{
		Path:  tripDto.ConnectPath(),
		Start: uint(start),
		End:   uint(end),
		Car:   uint(carId),
		Date:  date,
	}
	db.Create(&trip)
	err = db.Model(&cars[0]).Association("Trips").Append([]models.Trip{trip})
	cars[0].Miles = uint(end)
	db.Save(&cars[0])
	if err != nil {
		return nil, err
	}
	return &trip, nil
}

func GetTrips(idStr string, from string, to string) (interface{}, error) {
	db := database.GetDatabase()
	var err error
	fromDate, err := time.Parse(DateFormat, time.Time{}.Format(DateFormat))
	if err != nil {
		return nil, err
	}
	toDate, err := time.Parse(DateFormat, time.Now().Format(DateFormat))
	if len(from) != 0 {
		fromDate, err = time.Parse(DateFormat, from)
		if err != nil {
			return nil, err
		}
	}
	if len(to) != 0 {
		toDate, err = time.Parse(DateFormat, to)
		if err != nil {
			return nil, err
		}
	}
	if len(idStr) == 0 {
		var trips []models.Trip
		if len(to) == 0 {
			db.Where("date>=?", fromDate).Find(&trips)
		} else {
			db.Where("date>=? AND date<=?", fromDate, toDate).Find(&trips)
		}
		return trips, nil
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, err
	}
	var trips []models.Trip
	if len(to) == 0 {
		db.Where("date>=? AND id=?", fromDate, id).Find(&trips)
	} else {
		db.Where("date>=? AND date<=? AND id=?", fromDate, toDate, id).Find(&trips)
	}
	return trips[0], nil
}
