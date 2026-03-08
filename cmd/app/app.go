package main

import (
	"context"
	"log"
)

type App struct {
	*Container
}

func NewApp(ctx context.Context) (*App, error) {
	container, err := NewContainer(ctx)
	if err != nil {
		return nil, err
	}

	app := &App{
		Container: container,
	}

	log.Println("✅ App initialized")
	return app, nil
}

func (a *App) Run() error {
	url, err := a.Container.UrlShortnerService.CreateShortURL(context.Background(), "https://checkout.stripe.com/c/pay/cs_live_a1oRL1Ir8VSK8Rw5tZRpbAV4hzY8ryN8C23Yyg09C1ngLxFCYmFNBq7MP1#fidnandhYHdWcXxpYCc%2FJ2FgY2RwaXEnKSd2cGd2ZndsdXFsamtQa2x0cGBrYHZ2QGtkZ2lgYSc%2FY2RpdmApJ2R1bE5gfCc%2FJ3VuWmlsc2BaMDRJZzBJf0cxUV9SfVZMQlAxSWZkV3xzUTB9UjRMcj1fNkA0bEp1cEZnaUdJakltanQ2fVRrcTdQM3ZXNW5jdzNyclFhSWFURnxuY39La3IyMj1VZEoyazc1NXF2V2hTYDI3JyknY3dqaFZgd3Ngdyc%2FcXdwYCknZ2RmbmJ3anBrYUZqaWp3Jz8nJjU1NTU1NScpJ2lkfGpwcVF8dWAnPyd2bGtiaWBabHFgaCcpJ2BrZGdpYFVpZGZgbWppYWB3dic%2FcXdwYHgl")
	if err != nil {
		return err
	}
	log.Println("✅ Short URL created:", url)
	return nil
}
