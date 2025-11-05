package main

import (
	"backend-day1/middleware"
	"backend-day1/view"
	"fmt"
)

func main() {
	router := view.SetupRouter()
	router.Use(middleware.CORSMiddleware)
	fmt.Println("Server running")
	router.Run("localhost:8080")
}
