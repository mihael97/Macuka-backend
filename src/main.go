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
	for path, function := range carRoutes {
		r.HandleFunc(path, function).Methods("GET")
	}
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
