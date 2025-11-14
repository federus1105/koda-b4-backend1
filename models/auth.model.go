package models

import (
	"backend-day1/libs"
	"fmt"
)

type Auth struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name" binding:"max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"password_complex"`
}

var NextId = 1

func Register(a Auth) User {
	// --- HASHING PASSWORD ---
	hashedPassword, err := libs.HashPassword(a.Password)
	if err != nil {
		fmt.Println(err)
	}

	user := User{
		Id:       NextId,
		Name:     a.Name,
		Email:    a.Email,
		Password: hashedPassword,
	}

	NextId++
	Users = append(Users, user)
	return user
}

func Login(email, password string) (*User, string, bool) {
	// --- SEARCH USERS ---
	var storedUser User
	found := false
	for _, user := range Users {
		if user.Email == email {
			storedUser = user
			found = true
			break
		}
	}

	// --- IF USER NOT FOUND ---
	if !found {
		return nil, "Email atau password salah", false
	}

	// --- VERIFY PASSWORD ---
	match, err := libs.VerifyPassword(password, storedUser.Password)
	if err != nil || !match {
		return nil, "Email atau password salah", false
	}

	return &storedUser, "Login berhasil", true
}
