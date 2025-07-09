package configs

import "database/sql"

type Config struct {
	DB         *sql.DB
	AuthConfig *AuthConfig
	Handler    *Handler
	Env        string
}

type Handler struct{}

type AuthConfig struct {
	TTL       int
	SecretKey string
}
