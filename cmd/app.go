package main

import (
	"fmt"

	"github.com/alifrahmadian/habit-tracker-app-backend/configs"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/db"
	"github.com/alifrahmadian/habit-tracker-app-backend/internal/routes"
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

	return &configs.Config{
		DB:         db,
		AuthConfig: authConfig,
		Handler:    &configs.Handler{},
		Env:        env,
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
