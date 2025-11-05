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

func UpdateUser(id int, name, batch *string) *User {
	for i, u := range Users {
		if u.Id == id {
			if name != nil {
				Users[i].Name = *name
			}
			if batch != nil {
				Users[i].Batch = *batch
			}
			return &Users[i]
		}
	}
	return nil
}

func Login(email, password string) bool {
	for _, user := range UsersAuth {
		if user.Email == email && user.Password == password {
			return true
		}
	}
	return false
}
