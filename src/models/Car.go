package models

type Car struct {
	Id                int64  `json:"id" ,gorm:"primaryKey,autoIncrement"`
	Name              string `json:"name"`
	Miles             int64  `json:"miles"`
	RegistrationPlate string `json:"registrationPlate"`
	Year              uint   `json:"productionYear"`
	Trips             []Trip `gorm:"foreignKey:Car;references:Id;CASCADE:DELETE"`
}
