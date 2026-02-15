// Package config содержит функции для загрузки и валидации конфигурации приложения.
package config

import "github.com/spf13/viper"

// Config структура для хранения конфигурации приложения.
type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBPort int    `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
}

// LoadConfig загружает конфигурацию из файла и переменных окружения с помощью viper.
func LoadConfig(path string) (Config, error) {
	// Загрузка конфигурации из файла с помощью viper
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
