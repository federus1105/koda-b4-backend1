package controllers

import (
	"backend-day1/middleware"
	"backend-day1/models"
	"backend-day1/utils"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get All User
// @Summary      Get all user with pagination, search, and sorting
// @Description  Retrieve todos with optional search, pagination, and sort
// @Tags         User
// @Param        page        query  int    false  "Page number"         default(1)
// @Param        limit       query  int    false  "Items per page"      default(8)
// @Param        search      query  string false  "Search by name"
// @Param        sort_order  query  string false  "Sort order (ASC/DESC)" enums(ASC, DESC) default(ASC)
// @Success      200         {object}  models.ResponseSuccess
// @Router       /users [get]
// @Security BearerAuth
func GetAllUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "8"))
	search := ctx.DefaultQuery("search", "")
	sortOrder := ctx.DefaultQuery("sort_order", "ASC")

	users, msg := models.GetAllUsers(page, limit, search, sortOrder)
	if msg != "" {
		ctx.JSON(200, models.Response{
			Success: false,
			Message: msg,
		})
		return
	}

	// --- RESPONSE --
	var resp []models.UserResponse
	for _, u := range users {
		resp = append(resp, models.UserResponse{
			Id:      u.Id,
			Name:    u.Name,
			Email:   u.Email,
			Batch:   u.Batch,
			Profile: u.ProfileImages,
		})
	}
	ctx.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: "Get data successfully",
		Result:  resp,
	})
}

// GetUserbyID godoc
// @Summary      Get User by ID
// @Description  Retrieve a user item by its ID
// @Tags         User
// @Param        id   path      int  true  "user ID"
// @Success      200  {object}  models.ResponseSuccess
// @Router       /users/{id} [get]
// @Security BearerAuth
func GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(401, models.Response{
			Success: false,
			Message: "ID tidak valid",
		})
		return
	}
	user := models.GetUserById(id)
	if user == nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "User tidak ditemukan",
		})
		return
	}

	userResp := models.UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Batch:   user.Batch,
		Profile: user.ProfileImages,
	}

	ctx.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: "Get Data succesfully",
		Result:  userResp,
	})
}

// DeleteUser godoc
// @Summary      Delete a user by ID
// @Description  Delete User by its ID
// @Tags         User
// @Param        id  path      int  true  "User ID"
// @Success      200         {object}  models.ResponseSuccess
// @Router       /users/{id} [delete]
// @Security BearerAuth
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
	ctx.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: "User berhasil dihapus",
		Result: gin.H{
			"iduser": id,
		},
	})
}

// UpdateUser godoc
// @Summary      Edit an existing User by ID
// @Description  Update user's name, batch and profile images by its ID
// @Tags         User
// @Accept       multipart/form-data
// @Produce      json
// @Param        id            path      int     true   "User ID"
// @Param        name          formData  string  false   "User's new name"
// @Param        batch         formData  string  false   "User's new batch"
// @Param        profile 	   formData  file    false  "Profile image to upload"
// @Success      200     {object}  models.Response
// @Router       /users/{id} [patch]
// @Security BearerAuth
func UpdateUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(401, models.Response{
			Success: false,
			Message: "ID tidak valid"})
		return
	}
	var update models.UpdateUserRequest
	if err := ctx.ShouldBind(&update); err != nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: err.Error()})
		return
	}

	// --- CHECKING CLAIMS TOKEN ---
	claims, exists := ctx.Get("claims")
	if !exists {
		fmt.Println("ERROR :", !exists)
		ctx.AbortWithStatusJSON(403, models.Response{
			Success: false,
			Message: "Please log in again",
		})
		return
	}
	user, ok := claims.(middleware.Claims)
	if !ok {
		fmt.Println("ERROR", !ok)
		ctx.AbortWithStatusJSON(500, models.Response{
			Success: false,
			Message: "An error occurred!, please try again.",
		})
		return
	}
	var filenamePtr *string

	if update.ProfileImages != nil {
		saveDir := "upload/profile"
		ctxWithTimeout := context.Background()
		savePath, generatedFilename, err := utils.UploadImageFile(ctxWithTimeout, update.ProfileImages, saveDir, fmt.Sprintf("user_%d", user.ID))
		if err != nil {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if err := ctx.SaveUploadedFile(update.ProfileImages, savePath); err != nil {
			ctx.JSON(500, models.Response{
				Success: false,
				Message: "Gagal menyimpan file",
			})
			return
		}

		filenamePtr = &generatedFilename
	}
	updated := models.UpdateUser(userID, update.Name, update.Batch, filenamePtr)
	if updated == nil {
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "User tidak ditemukan"})
		return
	}
	// --- RESPONSE --
	userResp := models.UserResponse{
		Id:      updated.Id,
		Name:    updated.Name,
		Email:   updated.Email,
		Batch:   updated.Batch,
		Profile: updated.ProfileImages,
	}

	ctx.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: "update data successfully",
		Result:  userResp})
}

// UpdatePassword godoc
// @Summary      Update Password
// @Description  Update Password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param Auth body models.UpdatePasswordRequest true "Auth data"
// @Success      200   {object}  models.ResponseSuccess
// @Router       /users/update [post]
// @Security BearerAuth
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
		c.JSON(500, models.Response{
			Success: false,
			Message: msg,
		})
		return
	}

	var userResp any
	userResp = updatedUser
	userResp = models.UserResponse{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Batch:   user.Batch,
		Profile: user.ProfileImages,
	}

	c.JSON(200, models.ResponseSuccess{
		Success: true,
		Message: "Update password successfully",
		Result:  userResp,
	})
}
