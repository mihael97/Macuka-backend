package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"macuka-backend/src/database"
	"macuka-backend/src/dto"
	"macuka-backend/src/models"
	"macuka-backend/src/util"
	"net/http"
)

func CreateInvoice(invoiceDto dto.InvoiceDto) (models.Invoice, error) {
	db := database.GetDatabase()
	date, err := util.ConvertDate(invoiceDto.Date)
	if err != nil {
		return models.Invoice{}, err
	}
	currencyDate, err := util.ConvertDate(invoiceDto.CurrencyDate)
	if err != nil {
		return models.Invoice{}, err
	}

	invoice := models.Invoice{
		InvoiceNumber: invoiceDto.InvoiceNumber,
		Customer:      invoiceDto.Customer,
		Date:          *date,
		CurrencyDate:  *currencyDate,
		CallingNumber: invoiceDto.CallingNumber,
		HasVAT:        invoiceDto.HasVAT,
	}
	db.Create(&invoice)

	invoiceItems := make([]models.InvoiceItem, 0)
	for _, invoiceItemDto := range invoiceDto.InvoiceItems {
		invoiceItem := models.InvoiceItem{
			Description: invoiceItemDto.Description,
			Quantity:    invoiceItemDto.Quantity,
			Measure:     invoiceItemDto.Measure,
			Invoice:     invoice.Id,
		}
		db.Create(&invoiceItem)
		invoiceItems = append(invoiceItems, invoiceItem)
	}
	invoice.InvoiceItems = invoiceItems
	db.Updates(&invoice)

	log.Print("Added invoice {}", invoice)
	return invoice, nil
}

func ExportInvoice(id string, writer http.ResponseWriter) http.Header {
	db := database.GetDatabase()
	var invoices []models.Invoice
	db.Where(id).Find(&invoices)
	if len(invoices) == 0 {
		return nil
	}

	invoice := invoices[0]
	var customers []models.Customer
	db.Where("id=?", invoice.Customer).Find(&customers)
	customer := customers[0]

	data := make(map[string]interface{}, 0)
	data["invoice"] = invoice
	data["customer"] = customer
	data["creationDate"] = invoice.Date.Format("2.1.2006")
	data["currencyTime"] = invoice.CurrencyDate.Format("2.1.2006")

	var invoiceItems []models.InvoiceItem
	db.Raw("SELECT * FROM invoice_items WHERE invoice=?", invoice.Id).Find(&invoiceItems)
	data["invoiceItems"] = invoiceItems

	request, err := http.Get(fmt.Sprintf("%s/templates/generate", util.GetEnvVariable("DOCUMENT_API_URL", "https://document-creator.herokuapp.com")))
	if err != nil {
		return nil
	}
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil
	}
	writer.Write(content)
	return request.Header
}
