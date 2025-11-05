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
		user.POST("", controllers.Register)
		user.DELETE("/:id", controllers.DeleteUser)
		user.GET("/:id", controllers.GetUserById)
		user.PATCH("/:id", controllers.UpdateUserById)
		user.POST("/login", controllers.Login)
	}
	return router
}
