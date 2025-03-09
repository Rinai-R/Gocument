package Client

import (
	"context"
	pb "github.com/Rinai-R/Gocument/Server/Api/Func/Auth/Client/rpc"
	"github.com/Rinai-R/Gocument/Server/Api/Initialize"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	AuthClient pb.AuthClient
	AuthConn   *grpc.ClientConn
	PublicKey  string
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
	PublicKey = res.PublicKey
}
