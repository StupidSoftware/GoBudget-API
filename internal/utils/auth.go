package utils

import (
	"time"

	"github.com/breno5g/GoBudget/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string) (string, error) {
	token := config.GetTokenAuth()
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(config.GetEnv().JWTSecretKey))
}

func DecodeToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().JWTSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims["user_id"].(string), nil
}
