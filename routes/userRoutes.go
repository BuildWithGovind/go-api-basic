package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/", controllers.GetUsers)
	}
}
