package Client

import (
	"fmt"
	"github.com/Rinai-R/Gocument/App/Func/User/Client/rpc"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Registry"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient pb.UserClient
	UserConn   *grpc.ClientConn
)

func init() {

	param := vo.GetServiceParam{
		Clusters:    []string{"cluster1"},
		ServiceName: "User",
		GroupName:   "Gocument",
	}
	service, err := Registry.Client.GetService(param)
	if err != nil {
		Logger.Logger.Panic("Client " + err.Error())
	}

	if len(service.Hosts) == 0 {
		Logger.Logger.Panic("Client: service host is empty")
	}

	addr := fmt.Sprintf("%s:%d", service.Hosts[0].Ip, service.Hosts[0].Port)

	opt := grpc.WithTransportCredentials(insecure.NewCredentials())

	UserConn, err = grpc.Dial(addr, opt)
	if err != nil {
		Logger.Logger.Panic("Client: " + err.Error())
	}
	UserClient = pb.NewUserClient(UserConn)

	Logger.Logger.Debug("Client: User client OK")
}
