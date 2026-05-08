package main

import (
	"log"

	"github.com/hisshihi/url-shortener/internal/app"
	"github.com/hisshihi/url-shortener/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = app.New(cfg).Run()
	if err != nil {
		log.Fatal(err)
	}
}
