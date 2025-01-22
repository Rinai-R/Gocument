package Middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var UserSigningKey = []byte("114514")

func GenerateJWT(Message string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["GetName"] = Message
	claims["exp"] = time.Now().Add(60 * time.Hour).Unix()
	TokenString, err := token.SignedString(UserSigningKey)
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

func VerifyJWT(TokenString string) (string, error) {
	token, err := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return UserSigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok1 := token.Claims.(jwt.MapClaims); ok1 {
		message := claims["GetName"].(string)
		return message, nil
	}
	return "", errors.New("token invalid")
}
