package main

import (
	"flag"
	"github.com/endr-i/promo-back-go/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to config file")
}

func main() {
	flag.Parse()

	app.Run(configPath)
}
