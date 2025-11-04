package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required,max=20"`
	Batch string `json:"batch" binding:"required,max=2"`
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

var Users []User

func main() {
	router := gin.Default()

	// --- POST ---
	router.POST("/users", func(ctx *gin.Context) {
		body := User{}

		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(500, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		Users = append(Users, body)
		ctx.JSON(200, gin.H{
			"succes": true,
			"data":   Users,
		})
	})


	router.Run("localhost:8080")
}
