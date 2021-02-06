package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/mihael97/go-utility/src/web"
	jwtPackage "macuka-backend/src/jwt"
	"macuka-backend/src/models"
	"macuka-backend/src/services"
	"net/http"
	"time"
)

func GetAuthenticationRoutes() map[PathMethodPair]func(w http.ResponseWriter, r *http.Request) {
	routes := make(map[PathMethodPair]func(w http.ResponseWriter, r *http.Request), 0)

	routes[PathMethodPair{
		Path:   "/login",
		Method: PostMethod,
	}] = login

	return routes
}

func login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if !services.CheckUser(user) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	configValue := services.GetConfig("EXP_TOKEN")

	if configValue == nil {
		web.WriteErrorMessage("No exp token value", w)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(jwtPackage.JwtToken{Token: tokenString})
}
