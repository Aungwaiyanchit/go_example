package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Email    string `json:"email" gorm:"primary_key"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JwtUser struct {
	jwt.RegisteredClaims
	Email string
}