package view

import (
	"backend-day1/controllers"
	"backend-day1/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
