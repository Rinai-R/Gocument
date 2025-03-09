package main

import (
	"fmt"
	Initialize "github.com/Rinai-R/Gocument/Server/Auth/Registry"
	"github.com/Rinai-R/Gocument/Server/Auth/handle"
	pb "github.com/Rinai-R/Gocument/Server/Auth/rpc"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10003")
	if err != nil {
		Logger.Logger.Panic(err.Error())
	}
	Initialize.InitEtcd()
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	AuthSrv := handle.NewAuthService()
	pb.RegisterAuthServer(grpcServer, AuthSrv)
	Initialize.EtcdRegistry.ServiceRegister("Auth", "127.0.0.1:10003")

	msg := fmt.Sprintf("grpc server listening at %v", listener.Addr())
	fmt.Println(msg)
	Logger.Logger.Debug(msg)

	if err = grpcServer.Serve(listener); err != nil {
		Logger.Logger.Panic(err.Error())
	}
	defer grpcServer.GracefulStop()
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)
}
