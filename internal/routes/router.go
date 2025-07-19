package routes

import (
	"github.com/alifrahmadian/habit-tracker-app-backend/configs"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(secretKey string, router *gin.Engine, handlers *configs.Handler) {
	publicRoutes := router.Group("")
	{
		publicRoutes.POST("/register", handlers.AuthHandler.Register)
		publicRoutes.POST("/login", handlers.AuthHandler.Login)
	}
}
