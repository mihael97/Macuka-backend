package controllers

import (
	"github.com/gorilla/mux"
	"gitlab.com/mihael97/go-utility/src/web"
	"macuka-backend/src/services"
	"net/http"
)

func GetCarPaths() map[PathMethodPair]func(http.ResponseWriter, *http.Request) {
	routes := make(map[PathMethodPair]func(http.ResponseWriter, *http.Request))
	routes[PathMethodPair{
		Path:   "/cars/{id}",
		Method: GetMethod,
	}] = getCars
	routes[PathMethodPair{
		Path:   "/cars",
		Method: GetMethod,
	}] = getCars
	routes[PathMethodPair{
		Path:   "/cars",
		Method: PostMethod,
	}] = createCar
	routes[PathMethodPair{
		Path:   "/cars/{id}",
		Method: DeleteMethod,
	}] = deleteCar
	return routes
}

func deleteCar(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	err := services.DeleteCar(params["id"])
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func createCar(writer http.ResponseWriter, request *http.Request) {
	params, err := web.ParseParams([]string{"registrationPlate", "miles", "name", "productionYear"}, request)
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	car, err := services.CreateCar(params)
	if err != nil {
		web.WriteError(err, writer)
		return
	}
	web.ParseToJson(car, writer, http.StatusCreated)
}

func getCars(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	cars, err := services.GetCars(params["id"])
	if err != nil {
		web.WriteError(err, writer)
	} else {
		web.ParseToJson(cars, writer, http.StatusOK)
	}
}
