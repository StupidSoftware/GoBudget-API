package config

import (
	"github.com/golang-jwt/jwt/v5"
)

func initJWT() *jwt.Token {
	return jwt.New(jwt.SigningMethodHS256)
}
