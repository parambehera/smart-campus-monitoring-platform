package routes

import (
	"device-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// Home & Health
	router.GET("/", handlers.HomeHandler)
	router.GET("/health", handlers.HealthHandler)

	// Device APIs
	router.POST("/devices", handlers.CreateDeviceHandler)

	router.GET("/devices", handlers.GetAllDevicesHandler)

	router.GET("/devices/:id", handlers.GetDeviceHandler)

	router.PUT("/devices/:id", handlers.UpdateDeviceHandler)

	router.DELETE("/devices/:id", handlers.DeleteDeviceHandler)
}