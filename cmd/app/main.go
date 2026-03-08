// Package main точка входа в приложение.
package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	application, err := NewApp(ctx)
	if err != nil {
		log.Fatal("❌ Failed to create app:", err)
	}

	if err := application.Run(); err != nil {
		log.Fatal("❌ Failed to run app:", err)
	}
	log.Println("✅ Application started")
}
