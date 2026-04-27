// Package config содержит функции для загрузки и валидации конфигурации приложения.
package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Config структура для хранения конфигурации приложения.
type Config struct {
	DBHost   string `env:"DB_HOST"     envDefault:"localhost"`
	DBPort   int    `env:"DB_PORT"     envDefault:"5432"`
	DBUser   string `env:"DB_USER,required"`
	DBPass   string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8080"`
}

func (c Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName,
	)
}

func Load() (Config, error) {
	var cfg Config
	return cfg, env.Parse(&cfg)
}
