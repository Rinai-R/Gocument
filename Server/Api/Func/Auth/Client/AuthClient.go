package Client

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	pb "github.com/Rinai-R/Gocument/Server/Api/Func/Auth/Client/rpc"
	"github.com/Rinai-R/Gocument/Server/Api/Initialize"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	AuthClient pb.AuthClient
	AuthConn   *grpc.ClientConn
	PublicKey  *rsa.PublicKey
)

func init() {
	err := Initialize.ETCD.DiscoverService("Auth")
	if err != nil {
		Logger.Logger.Error("AuthClient: " + err.Error())
		return
	}
	addr := Initialize.ETCD.GetService("Auth")
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	AuthConn, err = grpc.Dial(addr, opts)
	if err != nil {
		Logger.Logger.Error("AuthClient: " + err.Error())
		return
	}
	AuthClient = pb.NewAuthClient(AuthConn)
	Logger.Logger.Debug("AuthClient: OK")
	res, err := AuthClient.GetPublicKey(context.Background(), &pb.GetPublicKeyRequest{})
	if err != nil {
		Logger.Logger.Error("AuthClient: " + err.Error())
		return
	}
	block, _ := pem.Decode([]byte(res.PublicKey))
	publicKeyInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	PublicKey = publicKeyInterface.(*rsa.PublicKey)
}
