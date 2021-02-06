package models

type AppConfig struct {
	Id    uint   `json:"id" ;gorm:"primaryKey,autoIncrement"`
	Name  string `json:"name" ;gorm:"unique"`
	Value string `json:"value"`
}
