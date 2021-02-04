package services

import (
	"macuka-backend/src/database"
	"macuka-backend/src/dto"
	"macuka-backend/src/models"
)

func CreateCustomer(model models.Customer) models.Customer {
	db := database.GetDatabase()
	db.Create(&model)
	return model
}

func GetCustomers() []models.Customer {
	db := database.GetDatabase()
	var customers []models.Customer
	db.Find(&customers)
	return customers
}

func GetCustomerPairs() []dto.CustomerPairDto {
	db := database.GetDatabase()
	var customers []models.Customer
	db.Find(&customers)
	returnCustomers := make([]dto.CustomerPairDto, 0)
	for _, customer := range customers {
		returnCustomers = append(returnCustomers, dto.CustomerPairDto{
			Id:   customer.Id,
			Iban: customer.Iban,
			Name: customer.Name,
		})
	}
	return returnCustomers
}
