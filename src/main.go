package main

import (
	"github.com/gorilla/mux"
	"log"
	"macuka-backend/src/controllers"
	"macuka-backend/src/database"
	"net/http"
)

func addRoute(r *mux.Router, routes map[controllers.PathMethodPair]func(w http.ResponseWriter, r *http.Request)) {
	for pathMethodPair, function := range routes {
		r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	}
}

func main() {
	database.InitializeDatabase()
	r := mux.NewRouter()
	//for pathMethodPair, function := range controllers.GetCarPaths() {
	//	r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	//}
	addRoute(r, controllers.GetCarPaths())
	//for pathMethodPair, function := range controllers.GetTripPaths() {
	//	r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	//}
	addRoute(r, controllers.GetTripPaths())
	//for pathMethodPair, function := range controllers.GetCustomerRoutes() {
	//	r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	//}
	addRoute(r, controllers.GetCustomerRoutes())
	//for pathMethodPair, function := range controllers.GetCityRoutes() {
	//	r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	//}
	addRoute(r, controllers.GetCityRoutes())

	//for pathMethodPair, function := range controllers.GetInvoiceRoutes() {
	//	r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	//}
	addRoute(r, controllers.GetInvoiceRoutes())

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
