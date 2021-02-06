package services

import (
	"macuka-backend/src/database"
	"macuka-backend/src/models"
)

func CheckUser(user models.User) bool {
	db := database.GetDatabase()
	var users []models.User
	db.Raw("SELECT * FROM users WHERE username=? AND password=?", user.Username, user.Password).Find(&users)
	return len(users) == 1
}
