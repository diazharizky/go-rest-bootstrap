package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
	})

	appSecret := os.Getenv("APP_JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(appSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyToken(tokenStr string) (bool, error) {
	appSecret := os.Getenv("APP_JWT_SECRET")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(appSecret), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
