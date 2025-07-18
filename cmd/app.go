package main

import (
	"fmt"

	"github.com/alifrahmadian/habit-tracker-app-backend/configs"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/db"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/handlers"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/repositories"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/routes"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *configs.Config
}

func loadConfig() (*configs.Config, error) {
	err := configs.LoadGoDotEnv()
	if err != nil {
		return nil, err
	}

	dbConfig := configs.LoadDBConfig()
	env := configs.LoadEnv()
	authConfig := configs.LoadAuthConfig()

	db, err := db.Connect(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	userRepo := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(&authService, authConfig.SecretKey, authConfig.TTL)

	return &configs.Config{
		DB:         db,
		AuthConfig: authConfig,
		Handler: &configs.Handler{
			AuthHandler: authHandler,
		},
		Env: env,
	}, nil
}

func NewApp() *App {
	cfg, err := loadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(cfg.AuthConfig.SecretKey, router, cfg.Handler)

	return &App{
		Router: router,
		Config: cfg,
	}
}
