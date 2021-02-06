package services

import (
	"macuka-backend/src/database"
	"macuka-backend/src/models"
)

func GetConfig(name string) *models.AppConfig {
	var config models.AppConfig
	db := database.GetDatabase()
	db.Where("NAME=?", name).Find(&config)
	if config.Name == "" && config.Value == "" {
		return nil
	}
	return &config
}
