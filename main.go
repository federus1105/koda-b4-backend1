package main

import (
	"backend-day1/view"
	"fmt"
)

func main() {
	router := view.SetupRouter()
	fmt.Println("Server running")
	router.Run("localhost:8080")
}
