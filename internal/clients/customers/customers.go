package customers

import (
	"fmt"
	"net/http"

	github.com/codevibe-de/goadv--orders/internal/model
)

const customersURL = "http://localhost:8081/"

func GetCustomerByPhoneNumber(phoneNumber string) model.Customer {
	requestURL := fmt.Sprintf(customersURL + "customer/" + phoneNumber, serverPort)
	res, err := http.Get(requestURL)

	c := model.Customer{}
	json.Decoder(res.Body).Decode(&c)

	return &c
}
