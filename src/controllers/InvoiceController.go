package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

	routes[PathMethodPair{
		Path:   "/invoices/export/{id}",
		Method: GetMethod,
	}] = exportInvoices

	return routes
}

func exportInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	name := services.ExportInvoice(mux.Vars(r)["id"], w)
	if len(name) == 0 {
		web.WriteErrorMessage("error during invoice exporting", w)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+name)
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
