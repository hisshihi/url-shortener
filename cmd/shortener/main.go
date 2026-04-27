package main

import (
	"log"

	"github.com/hisshihi/url-shortener/internal/app"
	"github.com/hisshihi/url-shortener/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	var cfg config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal(err)
	}
	a := app.New(cfg)
	if err = a.Run(); err != nil {
		log.Fatal(err)
	}
}
