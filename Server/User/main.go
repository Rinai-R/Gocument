package main

import (
	"github.com/Rinai-R/Gocument/Registry"
	pb "github.com/Rinai-R/Gocument/Server/User/rpc"
	"github.com/Rinai-R/Gocument/Server/User/service"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	pb.RegisterUserServer(grpcServer, &service.UserServer{})

	Registry.RegisterServiceInstance(Registry.Client, vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        10001,
		ServiceName: "User",
		GroupName:   "Gocument",
		ClusterName: "cluster1",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})

	log.Printf("grpc server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	defer grpcServer.GracefulStop()
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)
	defer Registry.DeRegisterServiceInstance(Registry.Client, vo.DeregisterInstanceParam{
		Ip:          "127.0.0.1", // 根据实际情况填写
		Port:        10001,       // gRPC服务的端口
		Cluster:     "cluster1",
		ServiceName: "User",
		GroupName:   "Gocument",
	})

}
