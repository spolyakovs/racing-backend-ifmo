package apiserver

import "github.com/spolyakovs/racing-backend-ifmo/internal/store"

type Config struct {
	BindAddr string `toml:"bind_addr"` // address of the server
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
