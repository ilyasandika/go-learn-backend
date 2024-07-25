package config

import "github.com/golang-jwt/jwt/v5"

var SecretKey = []byte("secret")

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
