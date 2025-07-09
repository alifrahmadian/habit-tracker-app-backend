package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alifrahmadian/habit-tracker-app-backend/internal/db"
	"github.com/joho/godotenv"
)

func LoadDBConfig() *db.Config {
	return &db.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func LoadAuthConfig() *AuthConfig {
	ttl, _ := strconv.Atoi(os.Getenv("TTL"))

	return &AuthConfig{
		TTL:       ttl,
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}

func LoadEnv() string {
	return os.Getenv("ENV")
}

func LoadGoDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	return nil
}
