package main

import (
	"context"
	"log"

	"github.com/hisshihi/url-shortener/core/service"
)

type Container struct {
	UrlShortnerService *service.URLShortnerService
}

func NewContainer(ctx context.Context) (*Container, error) {
	c := &Container{}
	
	c.UrlShortnerService = service.NewURLShortnerService()

	log.Println("✅ Container initialized")
	return c, nil
}
