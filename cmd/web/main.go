package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/codevibe-de/goadv--orders/internal/api"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP Network Port")
	flag.Parse()

	logHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(logHandler)

	// Setup Server
	logger.Info("Server starting", "addr", *addr)
	err := http.ListenAndServe(*addr, api.Routes(logger))
	logger.Error("Encountered unrecoverable Server Error", "err", err.Error())
	os.Exit(1)
}
