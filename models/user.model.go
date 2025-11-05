package models

import "github.com/go-playground/validator/v10"

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=20"`
}

type Auth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=20"`
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

func Register(u User) []User {
	u.Id = NextId
	NextId++
	Users = append(Users, u)
	return Users
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

func UpdateUser(id int, email, password *string) *User {
	for i, u := range Users {
		if u.Id == id {
			if email != nil {
				Users[i].Email = *email
			}
			if password != nil {
				Users[i].Password = *password
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
			case "Email":
				if e.Tag() == "required" {
					return false, "Email harus diisi"
				}
				if e.Tag() == "email" {
					return false, "Email harus sesuai format"
				}
			case "Password":
				if e.Tag() == "required" {
					return false, "Password harus diisi"
				}
				if e.Tag() == "max" {
					return false, "Password maksimal 20 karakter"
				}
			}
		}
	}

	for _, user := range Users {
		if user.Email == email && user.Password == password {
			return true, "Login berhasil"
		}
	}

	return false, "Email atau password salah"
}
