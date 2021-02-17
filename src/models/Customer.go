package models

type Customer struct {
	Id           uint      `json:"id" ;gorm:"primaryKey,autoIncrement"`
	Name         string    `json:"name"`
	Iban         string    `json:"iban" ;gorm:"unique"`
	Oib          string    `json:"oib"`
	Address      string    `json:"address"`
	PostalNumber uint      `json:"postalNumber"`
	City         string    `json:"city"`
	Invoices     []Invoice `gorm:"foreignKey:Customer;references:Id;CASCADE:DELETE"`
}
