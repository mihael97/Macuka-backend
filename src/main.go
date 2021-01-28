package main

import (
	"github.com/gorilla/mux"
	"log"
	"macuka-backend/src/controllers"
	"macuka-backend/src/database"
	"net/http"
)

func main() {
	database.InitializeDatabase()
	r := mux.NewRouter()
	carRoutes := controllers.GetCarPaths()
	for pathMethodPair, function := range carRoutes {
		r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	}
	for pathMethodPair, function := range controllers.GetTripPaths() {
		r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	}
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
