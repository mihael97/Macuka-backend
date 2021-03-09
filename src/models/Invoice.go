package models

import "time"

type Invoice struct {
	Id            uint          `json:"id" ;gorm:"primaryKey;autoIncrement"`
	InvoiceNumber string        `json:"invoiceNumber" ;gorm:"unique"`
	Customer      uint          `json:"customer"`
	Date          time.Time     `json:"date"`
	CurrencyDate  time.Time     `json:"currencyDate"`
	CallingNumber string        `json:"callingNumber"`
	HasVAT        bool          `json:"hasVat"`
	InvoiceItems  []InvoiceItem `gorm:"foreignKey:Invoice;references:Id;CASCADE:DELETE"`
}
