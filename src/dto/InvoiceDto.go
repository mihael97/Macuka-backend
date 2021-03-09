package dto

import "macuka-backend/src/models"

type InvoiceDto struct {
	Customer      uint                 `json:"customer"`
	InvoiceNumber string               `json:"invoiceNumber" ;gorm:"unique"`
	Date          string               `json:"date"`
	CurrencyDate  string               `json:"currencyDate"`
	CallingNumber string               `json:"callingNumber"`
	HasVAT        bool                 `json:"hasVat"`
	InvoiceItems  []models.InvoiceItem `json:"invoiceItems"`
}
