package fn

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/Api/Func/Auth/Client"
	pb "github.com/Rinai-R/Gocument/Server/Api/Func/Auth/Client/rpc"
	"github.com/Rinai-R/Gocument/pkg/Error"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	res, err := Client.AuthClient.GenerateToken(context.Background(), &pb.GenerateTokenRequest{Username: username})
	if err != nil {
		return "", err
	}
	return res.Token, nil
}

func ParseToken(tokenString string) (username string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Client.PublicKey, nil
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
