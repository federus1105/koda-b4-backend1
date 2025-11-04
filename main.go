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

type updateUser struct {
	Id    *int    `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Batch *string `json:"batch,omitempty"`
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
			"success": true,
			"data":    Users,
		})
	})

	// --- GET ALL  ---
	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"data":    Users,
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

	// --- DELETE USER BY ID ---
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
					"success": true,
					"message": "berhasil menghapus user",
				})
				return
			}
		}
	})

	// --- UPDATE USER BY ID ---
	router.PATCH("/users/:id", func(ctx *gin.Context) {
		var updateUser updateUser
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "ID tidak valid",
			})
			return
		}

		if err := ctx.ShouldBindBodyWithJSON(&updateUser); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		for i, user := range Users {
			if user.Id == id {
				if updateUser.Name != nil {
					Users[i].Name = *updateUser.Name
				}
				if updateUser.Batch != nil {
					Users[i].Batch = *updateUser.Batch
				}
				ctx.JSON(200, gin.H{
					"success": true,
					"data":    Users[i],
				})
				return
			}

		}
	})

	router.Run("localhost:8080")
}
