package dto

type InvoiceDto struct {
	Customer    uint    `json:"customer"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}
