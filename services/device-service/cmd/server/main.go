package main

import (
	"device-service/internal/config"
	"device-service/internal/kafka"
	"device-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	config.ConnectRedis()

	kafka.InitProducer()

	defer kafka.CloseProducer()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":" + config.GetEnv("SERVER_PORT"))
}