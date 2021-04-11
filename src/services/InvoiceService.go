package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"macuka-backend/src/database"
	"macuka-backend/src/dto"
	"macuka-backend/src/models"
	"macuka-backend/src/util"
	"net/http"
	"os"
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

func ExportInvoice(id string, writer http.ResponseWriter) (bool, error) {
	db := database.GetDatabase()
	var invoices []models.Invoice
	db.Where(id).Find(&invoices)
	if len(invoices) == 0 {
		return false, nil
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

	body := &bytes.Buffer{}
	jsonContent := make(map[string]interface{}, 0)
	jsonContent["data"] = data
	jsonContent["type"] = "docx"
	templateName := os.Getenv("TEMPLATE_NAME")
	if len(templateName) == 0 {
		templateName = "template"
	}
	jsonContent["name"] = templateName
	content, err := json.Marshal(jsonContent)
	fmt.Println(string(content))
	if err != nil {
		return false, err
	}
	body.Write(content)

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/templates/generate", util.GetEnvVariable("DOCUMENT_API_URL", "https://document-creator.herokuapp.com")), body)
	if err != nil {
		return false, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		return false, err
	}
	content, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	writer.Write(content)
	writer.Header().Set("Content-Disposition", response.Header.Get("Content-Disposition"))
	writer.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	writer.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	return true, nil
}
