package controllers

import (
	"backend-day1/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Get All User
// @Summary      Get all user with pagination, search, and sorting
// @Description  Retrieve todos with optional search, pagination, and sort
// @Tags         User
// @Param        page        query  int    false  "Page number"         default(1)
// @Param        limit       query  int    false  "Items per page"      default(8)
// @Param        search      query  string false  "Search by name"
// @Param        sort_order  query  string false  "Sort order (ASC/DESC)" enums(ASC, DESC) default(ASC)
// @Success      200         {object}  models.User
// @Failure      500         {object}  map[string]interface{}
// @Router       /users [get]
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

// GetUserbyID godoc
// @Summary      Get User by ID
// @Description  Retrieve a user item by its ID
// @Tags         User
// @Param        id   path      int  true  "user ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users/{id} [get]
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

// DeleteUser godoc
// @Summary      Delete a user by ID
// @Description  Delete User by its ID
// @Tags         User
// @Param        id  path      int  true  "User ID"
// @Success      200         {object}  map[string]interface{}
// @Failure      400         {object}  map[string]interface{}
// @Failure      500         {object}  map[string]interface{}
// @Router       /users/{id} [delete]
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

// UpdateUser godoc
// @Summary      Edit an existing User by ID
// @Description  Update user's name and batch by its ID
// @Tags         User
// @Param        id      path      int             true  "User ID"
// @Param        Users body        UpdateUser true  "Users data"
// @Success      200     {object}  UpdateUser
// @Failure      400     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /users/{id} [patch]
func UpdateUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "ID tidak valid"})
		return
	}
	var update models.UpdateUserRequest
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

	ctx.JSON(401, models.Response{
		Success: false,
		Message: "Email atau password salah",
	})
}

// UpdatePassword godoc
// @Summary      Update Password
// @Description  Update Password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param Auth body models.UpdatePasswordRequest true "Auth data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /users/update [post]
func UpdatePassword(c *gin.Context) {
	var req models.UpdatePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, models.Response{
			Success: false,
			Message: "Invalid request" + err.Error(),
		})
		return
	}

	var user models.User
	found := false
	for _, u := range models.Users {
		if u.Email == req.Email {
			user = u
			found = true
			break
		}
	}

	if !found {
		c.JSON(404, models.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	updatedUser, msg, err := models.UpdatePassword(user, req.NewPassword)
	if err != nil {
		c.JSON(500, gin.H{
			"message": msg,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"user":    updatedUser,
	})
}
