package routes

import (
	"github.com/car-rental-template/api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, health *controllers.HealthController) {
	router.GET("/health", health.GetStatus)
	api := router.Group("/api/v1")
	api.GET("/health", health.GetStatus)
}
