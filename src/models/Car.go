package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Id                int64  `json:"id"`
	Miles             int64  `json:"miles"`
	RegistrationPlate string `json:"registration_plate"`
}
