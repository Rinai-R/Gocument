package main

import (
	"fmt"
	Initialize "github.com/Rinai-R/Gocument/Server/User/Registry"
	"github.com/Rinai-R/Gocument/Server/User/handle"
	pb "github.com/Rinai-R/Gocument/Server/User/rpc"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	Initialize.InitEtcd()
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterUserServer(grpcServer, &handle.UserServer{})

	Initialize.EtcdRegistry.ServiceRegister("User", "127.0.0.1:10001")

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
