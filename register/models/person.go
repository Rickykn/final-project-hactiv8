package models

type Person struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Number_phone int    `json:"number_phone"`
}
