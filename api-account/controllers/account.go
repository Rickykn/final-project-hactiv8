package controllers

import (
	"api-account/database"
	"api-account/models"

	"fmt"
)

func CreateAccount(userInput *models.Account) *models.Account {
	db := database.Get()
	status := "Active"
	keyNumber := 900
	accountNumber := fmt.Sprintf("%d%d", keyNumber, userInput.Number_phone)
	newAccount := &models.Account{
		Email:          userInput.Email,
		Name:           userInput.Name,
		Account_number: accountNumber,
		Number_phone:   userInput.Number_phone,
		Status:         status,
	}
	fmt.Println(newAccount)

	db.Create(&newAccount)
	return newAccount
}
