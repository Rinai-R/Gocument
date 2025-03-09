package handle

import (
	"context"
	"crypto/rsa"
	pb "github.com/Rinai-R/Gocument/Server/Auth/rpc"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	PrivateKey *rsa.PrivateKey
	PublicKey  string
}

var Authsrv *AuthService

func (a *AuthService) GenerateToken(_ context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": request.Username,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(a.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateTokenResponse{Token: tokenString}, nil
}

func (a *AuthService) GetPublicKey(_ context.Context, _ *pb.GetPublicKeyRequest) (*pb.GetPublicKeyResponse, error) {
	return &pb.GetPublicKeyResponse{
		PublicKey: a.PublicKey,
	}, nil
}
