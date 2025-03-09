package handle

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	pb "github.com/Rinai-R/Gocument/Server/Auth/rpc"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	privateKey *rsa.PrivateKey
	publicKey  string
}

func (a *AuthService) GenerateToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": request.Username,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(a.privateKey)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateTokenResponse{Token: tokenString}, nil
}

func (a *AuthService) GetPublicKey(ctx context.Context, request *pb.GetPublicKeyRequest) (*pb.GetPublicKeyResponse, error) {
	return &pb.GetPublicKeyResponse{
		PublicKey: a.publicKey,
	}, nil
}

func NewAuthService() *AuthService {
	privatePEM, _ := os.ReadFile("keys/private.pem")
	block, _ := pem.Decode(privatePEM)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	publicPEM, _ := os.ReadFile("keys/public.pem")

	return &AuthService{
		privateKey: privateKey,
		publicKey:  string(publicPEM),
	}
}
