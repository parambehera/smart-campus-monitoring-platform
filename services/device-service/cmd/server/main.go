package main

import (
	"device-service/internal/config"
	"device-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	config.ConnectRedis()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":" + config.GetEnv("SERVER_PORT"))
}