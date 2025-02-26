package customers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/codevibe-de/goadv--orders/internal/model"
)

const (
	customersURL = "http://localhost:8081/"
	serverPort   = ":8081"
)

func GetCustomerByPhoneNumber(phoneNumber string, logger *slog.Logger) (*model.Customer, error) {
	requestURL := fmt.Sprintf(customersURL+"customer/"+phoneNumber, serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		logger.Error("Failed to get customer", "err", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	c := model.Customer{}
	body, err := io.ReadAll(res.Body)
	json.Unmarshal(body, &c)

	return &c, nil
}
