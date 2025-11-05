package controllers

import (
	"backend-day1/models"

	"github.com/gin-gonic/gin"
)

type UpdateUser struct {
	Name  *string `json:"name,omitempty"`
	Batch *string `json:"batch,omitempty"`
}

func GetAllUsers(ctx *gin.Context) {
	users, msg := models.GetAllUsers()
	if msg != "" {
		ctx.JSON(200, models.Response{
			Success: false,
			Message: msg,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"success": true,
		"data":    users,
	})
}
