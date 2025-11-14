package utils

type RegisterRequest struct {
	Name     string `example:"yourname"`
	Email    string `example:"youremail@gmail.com"`
	Password string `example:"Password!#"`
}

type LoginRequest struct {
	Email    string `example:"youremail@gmail.com"`
	Password string `example:"Password!#"`
}
