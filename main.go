package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required,max=20"`
	Batch string `json:"batch" binding:"required,max=2"`
}

type Response struct {
	Success bool   `json:"success"`
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

	// --- GET ALL  ---
	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"succes": true,
			"data":   Users,
		})
	})

	// --- GET USER BY ID ---
	router.GET("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "ID tidak valid",
			})
			return
		}
		for _, user := range Users {
			if user.Id == id {
				ctx.JSON(200, gin.H{
					"success": true,
					"user":    user,
				})
				return
			}
		}

	})

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "ID tidak valid",
			})
			return
		}

		for i, user := range Users {
			if user.Id == id {
				Users = append(Users[:i], Users[i+1:]...)
				ctx.JSON(200, gin.H{
					"succes":  true,
					"message": "berhasil menghapus data user",
				})
				return
			}
		}
	})

	router.Run("localhost:8080")
}
