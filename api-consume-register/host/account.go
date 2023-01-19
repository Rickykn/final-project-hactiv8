package host

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Account struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Number_phone int    `json:"number_phone"`
}

func SendCreateAccount(account Account) {

	var url = "http://localhost:8084/accounts"

	client := resty.New()
	respo, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(account).
		Post(url)

	if err == nil {
		if respo.StatusCode() == http.StatusCreated {
			// fmt.Println("Create Account success")
			return
		}
	}
	fmt.Println("Create Account Failed ")
}
