package Middleware

import (
	"fmt"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error"
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
		Logger.Logger.Debug(err.Error())
		return "", err
	}
	Logger.Logger.Debug("产生token成功")
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
		Logger.Logger.Debug(err.Error())
		return "", err
	}
	if claims, ok1 := token.Claims.(jwt.MapClaims); ok1 {
		message := claims["GetName"].(string)
		Logger.Logger.Debug("解析token成功")
		return message, nil
	}
	return "", Error.TokenExpired
}
