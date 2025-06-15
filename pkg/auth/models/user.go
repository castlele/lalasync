package models

import "github.com/golang-jwt/jwt/v5"

type UserLogin struct {
	Name     string
	Password string
	JWT      string
}

type UserClaims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}
