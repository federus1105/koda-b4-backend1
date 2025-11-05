package view

import (
	"backend-day1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/users")
	{
		user.GET("", controllers.GetAllUsers)
	}
	return router
}
