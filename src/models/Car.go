package models

type Car struct {
	Id                int64  `json:"id",gorm:"primaryKey,autoIncrement"`
	Miles             int64  `json:"miles"`
	RegistrationPlate string `json:"registration_plate"`
	Trips             []Trip `gorm:"foreignKey:Id;references:Id"`
}
