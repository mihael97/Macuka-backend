package models

type InvoiceItem struct {
	Id          uint    `json:"id" ;gorm:"primaryKey,autoIncrement"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Measure     string  `json:"measure"`
	Invoice     uint    `json:"invoice"`
}
