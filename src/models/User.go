package models

type User struct {
	Id       uint   `json:"id" ;gorm:"primaryKey,autoIncrement"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
