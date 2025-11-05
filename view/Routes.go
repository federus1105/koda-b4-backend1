package view

import (
	"backend-day1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/users")
	{
		user.POST("", controllers.Register)
		user.POST("/login", controllers.Login)
		user.POST("/update", controllers.UpdatePassword)
		user.GET("", controllers.GetAllUsers)
		user.DELETE("/:id", controllers.DeleteUser)
		user.GET("/:id", controllers.GetUserById)
		user.PATCH("/:id", controllers.UpdateUserById)
	}
	return router
}
