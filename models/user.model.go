package models

import (
	"fmt"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
	"github.com/matthewhartstonge/argon2"
)

type User struct {
	Id            int    `json:"id,omitempty"`
	Email         string `json:"email,omitempty" binding:"required,email"`
	Password      string `json:"password,omitempty" binding:"required,max=20"`
	Name          string `json:"name,omitempty" binding:"required,max=20"`
	Batch         string `json:"batch,omitempty" binding:"required,max=2"`
	ProfileImages string `json:"profile" form:"profile"`
}

type UpdatePasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

type Auth struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=20"`
}

type UpdateUserRequest struct {
	Name          *string               `json:"name,omitempty" form:"name"`
	Batch         *string               `json:"batch,omitempty" form:"batch"`
	ProfileImages *multipart.FileHeader `json:"profile" form:"profile"`
}

type Response struct {
	Success bool
	Message string
}

var Users []User
var NextId = 1
var validate = validator.New()

func GetAllUsers() ([]User, string) {
	if len(Users) == 0 {
		return nil, "List user kosong"
	}
	return Users, ""
}

func GetUserById(id int) *User {
	for _, user := range Users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func Register(u User) User {
	argon := argon2.DefaultConfig()

	hashedPassword, err := argon.HashEncoded([]byte(u.Password))
	if err != nil {
		fmt.Println(err)
	}
	u.Password = string(hashedPassword)

	u.Id = NextId
	NextId++
	Users = append(Users, u)
	return u
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

func Login(email, password string) (bool, string) {
	auth := Auth{
		Email:    email,
		Password: password,
	}

	err := validate.Struct(auth)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Field() {
			case "email":
				if e.Tag() == "required" {
					return false, "Email harus diisi"
				}
				if e.Tag() == "email" {
					return false, "Email harus sesuai format"
				}
			case "password":
				if e.Tag() == "required" {
					return false, "Password harus diisi"
				}
				if e.Tag() == "max" {
					return false, "Password maksimal 20 karakter"
				}
			}
		}
	}
	var storedUser User
	found := false
	for _, user := range Users {
		if user.Email == email {
			storedUser = user
			found = true
			break
		}
	}

	if !found {
		return false, "Email atau password salah"
	}

	match, err := argon2.VerifyEncoded([]byte(password), []byte(storedUser.Password))
	if err != nil {
		fmt.Printf("Error during verification: %v\n", err)
		return false, "Email atau password salah"
	}

	if !match {
		fmt.Println("invalid password attempt")
		return false, "Email atau password salah"
	}
	return true, "Login berhasil"
}
