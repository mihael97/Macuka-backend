package services

import (
	"macuka-backend/src/database"
	"macuka-backend/src/dto"
	"macuka-backend/src/models"
	"macuka-backend/src/util"
)

func CreateInvoice(invoiceDto dto.InvoiceDto) (models.Invoice, error) {
	db := database.GetDatabase()
	date, err := util.ConvertDate(invoiceDto.Date)
	if err != nil {
		return models.Invoice{}, err
	}
	invoice := models.Invoice{
		Customer: invoiceDto.Customer,
		Amount:   invoiceDto.Amount,
		Date:     *date,
	}
	db.Create(&invoice)
	return invoice, nil
}
