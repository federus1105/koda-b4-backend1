package main

import (
	"backend-day1/middleware"
	"backend-day1/routes"
	"fmt"

	"github.com/joho/godotenv"
)

// @title Auth
// @version 1.0
// @description This is a sample API minitask for backend.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token.
func main() {
	godotenv.Load()
	router := routes.SetupRouter()

	router.Use(middleware.CORSMiddleware)
	fmt.Println("Server running")
	router.Run("localhost:8080")
}
