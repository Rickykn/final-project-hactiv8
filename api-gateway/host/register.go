package host

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Person struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Number_phone int    `json:"number_phone"`
}

func RegisterGateway(person *Person) error {

	var url = "http://localhost:8081/register"

	client := resty.New()
	respo, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(person).
		Post(url)

	if err == nil {
		if respo.StatusCode() == http.StatusCreated {
			fmt.Println("Create Account success")

		}
		return nil
	}

	return err

}
