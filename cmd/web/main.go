package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/codevibe-de/goadv--orders/internal/api"
	"github.com/codevibe-de/goadv--orders/internal/config"
)

func main() {
	c := config.Config{}
	flag.StringVar(&c.OrdersAddr, "orders-addr", ":8080", "HTTP Addr for Orders Service")
	flag.StringVar(&c.CustomersAddr, "customer-addr", ":8181", "HTTP Addr for Customers Service")
	flag.Parse()

	logHandler := slog.NewTextHandler(os.Stdout, nil)
	c.Logger = slog.New(logHandler)

	// Setup Server
	c.Logger.Info("Server starting", "addr", c.OrdersAddr)
	err := http.ListenAndServe(c.OrdersAddr, api.Routes(&c))
	c.Logger.Error("Encountered unrecoverable Server Error", "err", err.Error())
	os.Exit(1)
}
