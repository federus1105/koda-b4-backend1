package models

import (
	"fmt"
	"mime/multipart"
	"sort"
	"strings"

	"github.com/matthewhartstonge/argon2"
)

type UpdatePasswordRequest struct {
	Email       string `json:"email" binding:"email"`
	NewPassword string `json:"new_password" binding:"password_complex"`
}

type User struct {
	Id            int    `json:"id,omitempty"`
	Email         string `json:"email,omitempty" binding:"required,email"`
	Password      string `json:"password,omitempty" binding:"assword_complex"`
	Name          string `json:"name,omitempty" binding:"required,max=20"`
	Batch         string `json:"batch,omitempty" binding:"gte=0"`
	ProfileImages string `json:"profile" form:"profile"`
}

type UpdateUserRequest struct {
	Name          *string               `json:"name,omitempty" form:"name" binding:"max=20"`
	Batch         *string               `json:"batch,omitempty" form:"batch" binding:"gte=0"`
	ProfileImages *multipart.FileHeader `json:"profile" form:"profile"`
}

var Users []User

func GetAllUsers(page, limit int, search string, sortOrder string) ([]User, string) {
	if len(Users) == 0 {
		return nil, "List user kosong"
	}

	// --- SEARCH --
	var filtered []User
	if search == "" {
		filtered = Users
	} else {
		for _, u := range Users {
			if strings.Contains(strings.ToLower(u.Name), strings.ToLower(search)) {
				filtered = append(filtered, u)
			}
		}
	}

	// --- SORTING ---
	sort.Slice(filtered, func(i, j int) bool {
		if strings.ToUpper(sortOrder) == "DESC" {
			return filtered[i].Name > filtered[j].Name
		}
		// --- DEFAULT ASC ---
		return filtered[i].Name < filtered[j].Name
	})

	// --- PAGINATION ---
	start := (page - 1) * limit
	if start > len(filtered)-1 {
		return []User{}, ""
	}
	end := start + limit
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], ""
}

func GetUserById(id int) *User {
	for _, user := range Users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func UpdatePassword(u User, newPassword string) (User, string, error) {
	argon := argon2.DefaultConfig()
	hashedPassword, err := argon.HashEncoded([]byte(newPassword))
	if err != nil {
		return u, "Failed to hash new password",
			fmt.Errorf("hashing error: %w", err)
	}
	for i := range Users {
		if Users[i].Email == u.Email {
			Users[i].Password = string(hashedPassword)
			return Users[i], "Password updated successfully", nil
		}
	}
	return User{},
		"User not found for update",
		fmt.Errorf("user not found")
}

func DeleteUser(id int) bool {
	for i, u := range Users {
		if u.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateUser(id int, name, batch, profileImage *string) *User {
	for i, u := range Users {
		if u.Id == id {
			if name != nil {
				Users[i].Name = *name
			}
			if batch != nil {
				Users[i].Batch = *batch
			}
			if profileImage != nil {
				Users[i].ProfileImages = *profileImage
			}
			return &Users[i]
		}
	}
	return nil
}
