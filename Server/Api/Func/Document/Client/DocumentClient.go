package Client

import (
	"github.com/Rinai-R/Gocument/Server/Api/Func/Document/Client/rpc"
	"github.com/Rinai-R/Gocument/Server/Api/Initialize"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	DocumentClient pb.DocumentClient
	DocumentConn   *grpc.ClientConn
	err            error
)

func init() {

	err = Initialize.ETCD.DiscoverService("Document")
	if err != nil {
		Logger.Logger.Error(err.Error())
		return
	}
	addr := Initialize.ETCD.GetService("Document")
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())

	DocumentConn, err = grpc.Dial(addr, opt)
	if err != nil {
		Logger.Logger.Panic("Client: " + err.Error())
	}
	DocumentClient = pb.NewDocumentClient(DocumentConn)
	Logger.Logger.Debug("Client: Document client OK")
}
