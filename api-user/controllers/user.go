package controllers

import (
	"api-user/database"
	"api-user/models"
)

func CreateUser(userInput *models.Person) *models.Person {
	db := database.Get()
	newUser := &models.Person{
		Email:        userInput.Email,
		Name:         userInput.Name,
		Address:      userInput.Address,
		Number_phone: userInput.Number_phone,
	}

	db.Create(newUser)
	return newUser
}
