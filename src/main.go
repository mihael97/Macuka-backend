package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"macuka-backend/src/controllers"
	"macuka-backend/src/database"
	"net/http"
	"os"
	"strings"
)

type Exception struct {
	Message string `json:"message"`
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(Exception{Message: err.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	}
}

func addRoute(r *mux.Router, routes map[controllers.PathMethodPair]func(w http.ResponseWriter, r *http.Request)) {
	for pathMethodPair, function := range routes {
		if pathMethodPair.Path == "/cars" {
			r.HandleFunc(pathMethodPair.Path, ValidateMiddleware(function)).Methods(pathMethodPair.GetMethod())
			continue
		}
		r.HandleFunc(pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	}
}

func main() {
	database.InitializeDatabase()
	r := mux.NewRouter()
	addRoute(r, controllers.GetCarPaths())
	addRoute(r, controllers.GetTripPaths())
	addRoute(r, controllers.GetCustomerRoutes())
	addRoute(r, controllers.GetCityRoutes())
	addRoute(r, controllers.GetInvoiceRoutes())
	addRoute(r, controllers.GetAuthenticationRoutes())

	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Print("Port is " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
