package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret  string
	BcryptSalt string
	Postgres   Postgres
}

type Postgres struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	Params   string
}

func New() *Config {
	_ = godotenv.Load()

	return &Config{
		JWTSecret:  os.Getenv("JWT_SECRET"),
		BcryptSalt: os.Getenv("BCRYPT_SALT"),
		Postgres: Postgres{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Params:   os.Getenv("DB_PARAMS"),
		},
	}
}
