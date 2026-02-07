// Package app хранит зависимости приложения и запуск приложения(lifecycle).
package app

import (
	"context"

	"gorm.io/gorm"
)

type Dependencies struct {
	// Инфраструктура
	DB *gorm.DB

	// Сервисы

}

func NewDependencies(ctx context.Context) (*Dependencies, error) {
	c := &Dependencies{}


	return c, nil
}