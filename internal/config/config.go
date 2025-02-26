package config

import "log/slog"

type Config struct {
	OrdersAddr    string
	CustomersAddr string
	Logger        *slog.Logger
}
