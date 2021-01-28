package controllers

import (
	"encoding/json"
	"gitlab.com/mihael97/go-utility/src/web"
	"macuka-backend/src/models"
	"macuka-backend/src/services"
	"net/http"
)

func GetTripPaths() map[PathMethodPair]func(http.ResponseWriter, *http.Request) {
	routes := make(map[PathMethodPair]func(http.ResponseWriter, *http.Request))
	routes[PathMethodPair{
		Path:   "/trips",
		Method: PostMethod,
	}] = createTrip
	return routes
}

func createTrip(writer http.ResponseWriter, request *http.Request) {
	var decoded models.TripDto
	err := json.NewDecoder(request.Body).Decode(&decoded)
	if err != nil {
		web.WriteError(err, writer)
	}
	trip, err := services.CreateTrip(decoded)
	if err != nil {
		web.WriteError(err, writer)
	}
	web.ParseToJson(trip, writer, http.StatusCreated)
}
