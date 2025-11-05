package controllers

import (
	"backend-day1/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "ID tidak valid",
		})
		return
	}
	user := models.GetUserById(id)
	if user == nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": "User tidak ditemukan",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"success": true,
		"data":    user,
	})
}

func Register(ctx *gin.Context) {
	var body models.User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(200, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	users := models.Register(body)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "ID tidak valid",
		})
		return
	}
	if !models.DeleteUser(id) {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "User tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User berhasil dihapus",
	})
}

func UpdateUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "ID tidak valid"})
		return
	}
	var update UpdateUser
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: err.Error()})
		return
	}
	updated := models.UpdateUser(id, update.Name, update.Batch)
	if updated == nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "User tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    updated})
}

func Login(ctx *gin.Context) {
	var body models.Auth
	if err := ctx.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, e := range ve {
				var msg string
				switch e.Field() {
				case "Email":
					switch e.Tag() {
					case "required":
						msg = "Email harus diisi"
					case "email":
						msg = "Email harus sesuai format"
					}
				case "Password":
					switch e.Tag() {
					case "required":
						msg = "Password harus diisi"
					case "max":
						msg = "Password maksimal 20 karakter"
					case "min":
						msg = "Password minimal 6 karakter"
					}
				}
				ctx.JSON(400, models.Response{
					Success: false,
					Message: msg,
				})
				return
			}
		}
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Format input tidak valid",
		})
		return
	}

	success, msg := models.Login(body.Email, body.Password)
	if success {
		ctx.JSON(200, gin.H{
			"success": true, 
			"message": msg})
		return
	}

	ctx.JSON(401, gin.H{
		"success": false,
		"message": "Email atau password salah",
	})
}
