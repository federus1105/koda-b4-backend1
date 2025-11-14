package models

type Response struct {
	Success bool
	Message string
}

type ResponseSuccess struct {
	Success bool
	Message string
	Result  any
}

type AuthResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Batch   string `json:"batch"`
	Profile string `json:"profile,omitempty"`
}
