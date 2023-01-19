package host

import (
	"api-consume-register/models"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func SendCreateUser(user models.Person) {

	var url = "http://localhost:8082/users"

	client := resty.New()
	respo, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		Post(url)

	if err == nil {
		if respo.StatusCode() == http.StatusCreated {
			// fmt.Println("Create user success")
			return
		}
	}
	fmt.Println("Create Account Failed ")
}
