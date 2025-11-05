package main

import (
	"backend-day1/middleware"
	"backend-day1/view"
	"fmt"
)

// @title Auth
// @version 1.0
// @description This is a sample API minitask for backend.
// @host localhost:8080
// @BasePath /
func main() {
	router := view.SetupRouter()
	router.Use(middleware.CORSMiddleware)
	fmt.Println("Server running")
	router.Run("localhost:8080")
}
