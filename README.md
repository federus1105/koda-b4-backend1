#  Auth Backend API
> This is a small authentication project to understand the basic structure of a simple backend API builder, with some features such as password hashing using argon2, jwt token for user authentication and authorization and there is a cors middleware to allow a domain/url to access this backend source.

 ## 📸 Preview
 ### Swagger Documentation
![swagger](/assets/swaggerdocumentation.png)

## Feature
- 🔐 JWT Authentication (Bearer Token)
- 🔍 Search / Filtering
- ↕ Sorting (ASC / DESC)
- 📄 Pagination
- 📘 Swagger UI Documentation
- ⚡ Built with Go + Gin

## 🛠️ Tech Stack
![Go](https://img.shields.io/badge/-Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![Gin](https://img.shields.io/badge/-Gin-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![Swagger](https://img.shields.io/badge/Swagger-UI-85EA2D?logo=swagger&logoColor=black&style=for-the-badge)

##  🔐 .env Configuration
```
JWT_SECRET=your_jwt_secret
ORIGIN_URL=your_url
```

## 📦 How to Install & Run Project
### First, clone this repository: 
```
https://github.com/federus1105/koda-b4-backend1.git
```
### Install Dependencies
```go
go mod tidy
```
### Run Server/Project
```go
go run main.go 
```
### Init Swagger
```go
swag init -g main.go
```
### Open Swagger Documentation in Browser
#### ⚠️ Make sure the server is running
```http://localhost:8080/swagger/index.html```


## 👨‍💻 Made with by
📫 [federusrudi@gmail.com](mailto:federusrudi@gmail.com)  
💼 [LinkedIn](https://www.linkedin.com/in/federus-rudi/)  


## 📜 License
Released under the **MIT License**.  
You’re free to use, modify, and distribute this project — just don’t forget to give a little credit

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

