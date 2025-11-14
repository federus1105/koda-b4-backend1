package routes

import (
	"backend-day1/controllers"
	"backend-day1/docs"
	"backend-day1/libs"
	"backend-day1/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	libs.InitValidator()

	// --- SWAGGER ---
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// --- AUTH ROUTER ---
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// --- USER ROUTER ---
	user := router.Group("/users")
	{
		user.POST("", middleware.VerifyToken, controllers.UpdatePassword)
		user.GET("", middleware.VerifyToken, controllers.GetAllUsers)
		user.DELETE("/:id", middleware.VerifyToken, controllers.DeleteUser)
		user.GET("/:id", middleware.VerifyToken, controllers.GetUserById)
		user.PATCH("/:id", middleware.VerifyToken, controllers.UpdateUserById)
	}

	return router
}
