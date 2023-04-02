package config

import "github.com/endr-i/promo-back-go/internal/connection/pg"

type Config struct {
	PG       pg.PGConfig
	Rest     RestServerConfig
	LogLevel string `default:"error"`
}

type RestServerConfig struct {
	Address string `default:":8000"`
}
