package controllers

import (
	"encoding/json"
	"gitlab.com/mihael97/go-utility/src/web"
	"macuka-backend/src/dto"
	"macuka-backend/src/services"
	"net/http"
)

func GetInvoiceRoutes() map[PathMethodPair]func(w http.ResponseWriter, r *http.Request) {
	routes := make(map[PathMethodPair]func(w http.ResponseWriter, r *http.Request), 0)

	routes[PathMethodPair{
		Path:   "/invoices",
		Method: PostMethod,
	}] = createInvoice

	return routes
}

func createInvoice(w http.ResponseWriter, r *http.Request) {
	var invoiceDto dto.InvoiceDto
	err := json.NewDecoder(r.Body).Decode(&invoiceDto)
	if err != nil {
		web.WriteError(err, w)
		return
	}
	invoice, err := services.CreateInvoice(invoiceDto)
	if err != nil {
		web.WriteError(err, w)
		return
	}
	web.ParseToJson(invoice, w, http.StatusCreated)
}
