package config

type Config struct {
	PG struct {
		ConnectionString string
	}
	Rest     RestServerConfig
	LogLevel string `default:"error"`
}

type RestServerConfig struct {
	Address string `default:":8000"`
}
