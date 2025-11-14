package controllers

import (
	"backend-day1/middleware"
	"backend-day1/models"
	"backend-day1/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Register godoc
// @Summary 	Register user
// @Description Register user and Hash Password
// @Tags 		Auth
// @Param 		Register body 	utils.RegisterRequest  true 	"Register Info"
// @Success 	200 {object} 	models.ResponseSuccess
// @Router 		/auth/register [post]
func Register(ctx *gin.Context) {
	var input models.Auth
	// --- VALIDATION ---
	if err := ctx.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			var msgs []string
			for _, fe := range ve {
				msgs = append(msgs, utils.ErrorMessage(fe))
			}
			ctx.JSON(400, models.Response{
				Success: false,
				Message: strings.Join(msgs, ", "),
			})
			return
		}

		ctx.JSON(400, models.Response{
			Success: false,
			Message: "invalid JSON format",
		})
		return
	}

	user := models.Register(input)
	userResp := models.AuthResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: "Register Sucessfuly",
		Result:  userResp,
	})
}

// Login godoc
// @Summary 	Login user
// @Description Login user and get JWT token
// @Tags 		Auth
// @Param 		login body 		utils.LoginRequest  true 	"Login Info"
// @Success 	200 {object} 	models.ResponseSuccess
// @Router 		/auth/login [post]
func Login(ctx *gin.Context) {
	var input models.Auth

	// --- VALIDATION ---
	if err := ctx.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			var msgs []string
			for _, fe := range ve {
				msgs = append(msgs, utils.ErrorMessage(fe))
			}
			ctx.JSON(400, models.Response{
				Success: false,
				Message: strings.Join(msgs, ", "),
			})
			return
		}

		ctx.JSON(400, models.Response{
			Success: false,
			Message: "invalid JSON format",
		})
		return
	}

	// --- MODEL LOGIN ---
	user, msg, success := models.Login(input.Email, input.Password)
	if !success {
		ctx.JSON(401, models.Response{
			Success: false,
			Message: msg,
		})
		return
	}
	// --- GENERATE JWT TOKEN
	claims := middleware.NewJWTClaims(user.Id)
	jwtToken, err := claims.GenToken()
	if err != nil {
		fmt.Println("Internal Server Error.\nCause: ", err)
		ctx.JSON(500, models.Response{
			Success: false,
			Message: "internal server errorrr",
		})
		return
	}

	// --- RESPONSE --
	userResp := models.AuthResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: msg,
		Result: gin.H{
			"data":  userResp,
			"token": jwtToken,
		},
	})
}
