package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	DSN string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, using default config")
	}

	return &Config{
		Db: DbConfig{
			DSN: os.Getenv("DSN"),
		},

		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
