package api

import (
	"encoding/json"
	"net/http"

	"github.com/codevibe-de/goadv--orders/internal/config"
	"github.com/codevibe-de/goadv--orders/internal/model"
)

type ResponseData struct {
	Msg         string
	Params      map[string][]string
	Err         error
	CurrentTime string
}

func RequestLoggerMiddleware(c *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Logger.Info("Received request", "ip", r.RemoteAddr, "proto", r.Proto, "method", r.Method, "uri", r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func PlaceOrdersHandler(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o := model.Order{}
		json.NewDecoder(r.Body).Decode(&o)

		c.Logger.Info("Received order", "order", o)
	}
}

func Routes(c *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", PlaceOrdersHandler(c))

	return RequestLoggerMiddleware(c, mux)
}
