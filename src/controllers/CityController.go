package controllers

import (
	"encoding/json"
	"gitlab.com/mihael97/go-utility/src/web"
	"macuka-backend/src/models"
	"macuka-backend/src/services"
	"net/http"
)

func GetCityRoutes() map[PathMethodPair]func(writer http.ResponseWriter, request *http.Request) {
	routes := make(map[PathMethodPair]func(writer http.ResponseWriter, r *http.Request), 0)

	routes[PathMethodPair{
		Path:   "/cities",
		Method: PostMethod,
	}] = createRoutes

	routes[PathMethodPair{
		Path:   "/cities",
		Method: GetMethod,
	}] = getCities

	return routes
}

func getCities(writer http.ResponseWriter, r *http.Request) {
	web.ParseToJson(services.GetCities(), writer, http.StatusOK)
}

func createRoutes(writer http.ResponseWriter, r *http.Request) {
	var cities []models.City
	err := json.NewDecoder(r.Body).Decode(&cities)
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	services.CreateCities(cities)
	writer.WriteHeader(http.StatusCreated)
}
