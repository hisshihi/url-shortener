package main

import (
	"log"

	"github.com/hisshihi/url-shortener/core/app"
	"github.com/hisshihi/url-shortener/core/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	_ = app.New(cfg)
}
