package main

import (
	"go-api/config"
	"go-api/middleware"
	"go-api/models"
	"go-api/routes"
	"github.com/gin-gonic/gin"
)
func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})

	router := gin.Default()
	router.Use(middleware.Logger())

	routes.UserRoutes(router)

	router.Run(":8080")
}
