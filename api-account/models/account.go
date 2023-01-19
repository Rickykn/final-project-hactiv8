package models

type Account struct {
	ID             int    `json:"id" gorm:"primary_key"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Account_number string `json:"account_number"`
	Number_phone   int    `json:"number_phone"`
	Status         string `json:"status"`
}
