package Client

import (
	"github.com/Rinai-R/Gocument/Server/Api/Func/User/Client/rpc"
	"github.com/Rinai-R/Gocument/Server/Api/Initialize"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient pb.UserClient
	UserConn   *grpc.ClientConn
	err        error
)

func init() {

	err = Initialize.ETCD.DiscoverService("User")
	if err != nil {
		Logger.Logger.Error("Client: " + err.Error())
		return
	}
	addr := Initialize.ETCD.GetService("User")

	opt := grpc.WithTransportCredentials(insecure.NewCredentials())

	UserConn, err = grpc.Dial(addr, opt)
	if err != nil {
		Logger.Logger.Panic("Client: " + err.Error())
	}
	UserClient = pb.NewUserClient(UserConn)

	Logger.Logger.Debug("Client: User client OK")
}
