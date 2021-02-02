package models

type Car struct {
	Id                uint   `json:"id" ,gorm:"primaryKey,autoIncrement"`
	Name              string `json:"name"`
	Miles             uint   `json:"miles"`
	RegistrationPlate string `json:"registrationPlate"`
	Year              uint   `json:"productionYear"`
	Trips             []Trip `gorm:"foreignKey:Car;references:Id;CASCADE:DELETE"`
}
