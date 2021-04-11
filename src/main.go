package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"macuka-backend/src/controllers"
	"macuka-backend/src/database"
	"macuka-backend/src/dto"
	"macuka-backend/src/util"
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
			r.HandleFunc("/api"+pathMethodPair.Path, ValidateMiddleware(function)).Methods(pathMethodPair.GetMethod())
			continue
		}
		r.HandleFunc("/api"+pathMethodPair.Path, function).Methods(pathMethodPair.GetMethod())
	}
}

func initializeTemplate() {
	path := util.GetEnvVariable("DOCUMENT_API_URL", "https://document-creator.herokuapp.com") + "/documents"

	if !strings.Contains(path, "localhost") {
		return
	}

	file, err := ioutil.ReadFile("proba.xml")

	if err != nil {
		log.Fatal(err)
	}

	body := &bytes.Buffer{}
	content, err := json.Marshal(dto.DocumentDto{
		Name:    "template",
		Content: base64.StdEncoding.EncodeToString(file),
	})
	if err != nil {
		log.Fatal(err)
	}
	body.Write(content)
	request, err := http.NewRequest("POST", path, body)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(response.StatusCode)
}

func main() {
	database.InitializeDatabase()
	initializeTemplate()
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
