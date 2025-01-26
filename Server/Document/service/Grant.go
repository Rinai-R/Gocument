package service

import (
	"context"
	"github.com/Rinai-R/Gocument/DataBase/Document/dao"
	"github.com/Rinai-R/Gocument/Logger"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/models"
)

func (*DocumentServer) Grant(ctx context.Context, req *pb.GrantRequest) (*pb.GrantResponse, error) {
	//先判断请求的人的身份是不是文档主人
	if err := dao.IsHost(ctx, req.Host, int(req.DocumentId)); err != nil {
		Logger.Logger.Error("Grant Failed " + err.Error())
		return &pb.GrantResponse{
			Code: int64(ErrCode.GrantFailed),
			Msg:  err.Error(),
		}, nil
	}
	permission := models.Permission{
		DocumentId: int(req.DocumentId),
		UserId:     int(req.UserId),
		Type:       false,
	}

	if err := dao.Grant(permission); err != nil {
		return &pb.GrantResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}

	return &pb.GrantResponse{
		Code: int64(ErrCode.OK),
		Msg:  "ok",
	}, nil
}
