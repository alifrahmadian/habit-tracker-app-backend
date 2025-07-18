package configs

import (
	"database/sql"

	"github.com/alifrahmadian/habit-tracker-app-backend/internal/handlers"
)

type Config struct {
	DB         *sql.DB
	AuthConfig *AuthConfig
	Handler    *Handler
	Env        string
}

type Handler struct {
	AuthHandler *handlers.AuthHandler
}

type AuthConfig struct {
	TTL       int
	SecretKey string
}
