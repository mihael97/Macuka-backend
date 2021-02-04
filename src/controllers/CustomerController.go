package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/mihael97/go-utility/src/web"
	"macuka-backend/src/models"
	"macuka-backend/src/services"
	"net/http"
)

func GetCustomerRoutes() map[PathMethodPair]func(http.ResponseWriter, *http.Request) {
	routes := make(map[PathMethodPair]func(http.ResponseWriter, *http.Request))

	routes[PathMethodPair{
		Path:   "/customers",
		Method: GetMethod,
	}] = getCustomers
	routes[PathMethodPair{
		Path:   "/customers/{type}",
		Method: GetMethod,
	}] = getCustomers
	routes[PathMethodPair{
		Path:   "/customers",
		Method: PostMethod,
	}] = createCustomer

	return routes
}

func createCustomer(writer http.ResponseWriter, request *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(request.Body).Decode(&customer)
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	customer = services.CreateCustomer(customer)
	web.ParseToJson(customer, writer, http.StatusCreated)
}

func getCustomers(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	if params["type"] != "pairs" {
		web.ParseToJson(services.GetCustomers(), writer, http.StatusOK)
	} else {
		web.ParseToJson(services.GetCustomerPairs(), writer, http.StatusOK)
	}
}
