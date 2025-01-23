package service

import (
	"context"
	"github.com/Rinai-R/Gocument/DataBase/Document/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/models"
)

type DocumentServer struct {
	pb.UnimplementedDocumentServer
}

func (*DocumentServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	Document := models.Document{
		Username:  req.Username,
		Title:     req.Title,
		IsPrivate: req.IsPrivate,
	}

	if err := dao.Create(Document); err != nil {
		return &pb.CreateResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}
	return &pb.CreateResponse{
		Code: int64(ErrCode.OK),
		Msg:  "OK",
	}, nil
}
