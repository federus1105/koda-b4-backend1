package controllers

import (
	"backend-day1/models"
	"backend-day1/utils"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Register godoc
// @Summary      Register
// @Description  Register
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Auth  body      models.User  true  "Auth data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /users [post]
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
// @Summary      Login
// @Description  Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Auth  body      models.User  true  "Auth data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /users/login [post]
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

	// --- RESPONSE --
	userResp := models.AuthResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: msg,
		Result:  userResp,
	})
}
