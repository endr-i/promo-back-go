package config

type Config struct {
	PG       PGConfig
	Rest     RestServerConfig
	LogLevel string `default:"error"`
}

type RestServerConfig struct {
	Address string `default:":8000"`
}

type PGConfig struct {
	ConnectionString     string `default:"postgres://postgres:pgPass@localhost:5432/promo_db"`
	ConnectTimeout       int    `default:"5000"`
	PreferSimpleProtocol bool   `default:"true"`
	MaxOpenConns         int    `default:"100"`
	MaxIdleConns         int    `default:"1"`
}
