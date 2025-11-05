package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required,max=20"`
	Batch string `json:"batch" binding:"required,max=2"`
}

type Auth struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type Response struct {
	Success bool
	Message string
}

var UsersAuth = []Auth{
	{Email: "andi@example.com", Password: "password123"},
}

var Users []User
var NextId = 1

func GetAllUsers() ([]User, string) {
	if len(Users) == 0 {
		return nil, "List user kosong"
	}
	return Users, ""
}

