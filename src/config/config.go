package config

import (
	"dating-app/src/infrastructure/database"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
}

type App struct {
	Name    string
	Version string
	Port    string
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	if err = godotenv.Load(); err != nil {
		err = fmt.Errorf("error loading .env file: %w", err)
		return
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    os.Getenv("APP_PORT"),
		},
		Database: database.DatabaseConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
	}

	return
}
