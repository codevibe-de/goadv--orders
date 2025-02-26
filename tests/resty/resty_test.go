package resty_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/codevibe-de/goadv--orders/internal/api"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// #TODO:
// http://ordersvc.tld/placeOrder
// http://customer.tld/createCustomer
// http://customer.tld/getCustomers
// http://customer.tld/getCustomerByPhoneNumber

func TestPlaceOrdersHandler(t *testing.T) {
	// Create a new logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	// Create a new HTTP server using the Routes function
	server := httptest.NewServer(api.Routes(logger))
	defer server.Close()

	// Create a new Resty client
	client := resty.New()

	// Define a sample order payload
	order := map[string]interface{}{
		"product_id": 123,
		"quantity":   2,
	}

	// Convert the order to JSON
	orderJSON, err := json.Marshal(order)
	assert.NoError(t, err)

	// Send a POST request to the /products endpoint
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bytes.NewReader(orderJSON)).
		Post(server.URL + "/products")

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert the response status code is 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}
