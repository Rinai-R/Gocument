package handle

import (
	"context"
	dao2 "github.com/Rinai-R/Gocument/Server/Document/DataBase/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/pkg/Error/ErrCode"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/Rinai-R/Gocument/pkg/models"
)

func (*DocumentServer) Grant(ctx context.Context, req *pb.GrantRequest) (*pb.GrantResponse, error) {
	//先判断请求的人的身份是不是文档主人
	if err := dao2.IsHost(ctx, req.Host, int(req.DocumentId)); err != nil {
		Logger.Logger.Error("Grant Failed " + err.Error())
		return &pb.GrantResponse{
			Code: int64(ErrCode.GrantFailed),
			Msg:  err.Error(),
		}, nil
	}
	permission := models.Permission{
		DocumentId: int(req.DocumentId),
		UserId:     int(req.UserId),
		Type:       req.Type,
	}

	if err := dao2.Grant(ctx, permission); err != nil {
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
