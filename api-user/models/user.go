package models

type Person struct {
	ID           int    `json:"id" gorm:"primary_key"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Number_phone int    `json:"number_phone"`
}
