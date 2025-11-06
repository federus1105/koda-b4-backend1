package view

import (
	"backend-day1/controllers"
	"backend-day1/docs"
	"backend-day1/middleware"

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
		user.POST("/update", middleware.VerifyToken, controllers.UpdatePassword)
		user.GET("", middleware.VerifyToken, controllers.GetAllUsers)
		user.DELETE("/:id", middleware.VerifyToken, controllers.DeleteUser)
		user.GET("/:id", middleware.VerifyToken, controllers.GetUserById)
		user.PATCH("/:id", middleware.VerifyToken, controllers.UpdateUserById)
	}
	return router
}
