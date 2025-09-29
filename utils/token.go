package utils

import (
	"github.com/diazharizky/go-rest-bootstrap/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
	})

	appSecret := config.Global.GetString("app.jwt_secret")
	tokenStr, err := token.SignedString([]byte(appSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyToken(tokenStr string) (bool, error) {
	appSecret := config.Global.GetString("app.jwt_secret")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(appSecret), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
