package app

import (
	"github.com/endr-i/promo-back-go/internal/config"
	"github.com/endr-i/promo-back-go/internal/server"
	"github.com/jinzhu/configor"
	"log"
)

func Run(configPath string) {
	var cfg = &config.Config{}
	configor.New(&configor.Config{ENVPrefix: "PROMO"}).Load(cfg, configPath)

	srv := server.New(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}
}
