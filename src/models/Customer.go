package models

type Customer struct {
	Id            uint   `json:"id" ;gorm:"primaryKey,autoIncrement"`
	Name          string `json:"name"`
	Iban          string `json:"iban" ;gorm:"unique"`
	AccountNumber string `json:"account"`
	Address       string `json:"address"`
	PostalNumber  uint   `json:"postal_number"`
	City          string `json:"city"`
}
