package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetCarPaths() map[string]func(http.ResponseWriter, *http.Request) {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes["/cars/{id}"] = getCars
	return routes
}

func getCars(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	writer.Write([]byte(params["id"]))
}
