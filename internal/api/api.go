package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/codevibe-de/goadv--orders/internal/model"
)

type ResponseData struct {
	Msg         string
	Params      map[string][]string
	Err         error
	CurrentTime string
}

func RequestLoggerMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request", "ip", r.RemoteAddr, "proto", r.Proto, "method", r.Method, "uri", r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func PlaceOrdersHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o := model.Order{}
		json.NewDecoder(r.Body).Decode(&o)

		logger.Info("Received order", "order", o)
	}
}

func Routes(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", PlaceOrdersHandler(logger))

	return RequestLoggerMiddleware(logger, mux)
}
