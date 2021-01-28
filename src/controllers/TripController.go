package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	routes[PathMethodPair{
		Path:   "/trips",
		Method: GetMethod,
	}] = getTrips
	return routes
}

func getTrips(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	from := query.Get("from")
	to := query.Get("to")
	params := mux.Vars(request)
	trips, err := services.GetTrips(params["id"], from, to)
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	web.ParseToJson(trips, writer, http.StatusOK)
}

func createTrip(writer http.ResponseWriter, request *http.Request) {
	var decoded models.TripDto
	err := json.NewDecoder(request.Body).Decode(&decoded)
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	trip, err := services.CreateTrip(decoded)
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	web.ParseToJson(trip, writer, http.StatusCreated)
}
