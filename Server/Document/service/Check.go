package service

import (
	"context"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/models"
)

func (*DocumentServer) Check(c context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	id, err := dao.GetId(c, req.Username)
	if err != nil {
		return &pb.CheckPermissionResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}
	permission := models.Permission{
		DocumentId: int(req.DocumentId),
		UserId:     id,
	}
	err, msg := dao.Check(c, permission)
	if err != nil {
		if msg == "notfound" {
			Logger.Logger.Debug("Dao: Not Found")
			return &pb.CheckPermissionResponse{
				Code: int64(ErrCode.DocumentNotFound),
				Msg:  Error.DocumentNotFound.Error(),
			}, nil
		}
		Logger.Logger.Debug("Dao: Internal Error" + err.Error())
		return &pb.CheckPermissionResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}

	Logger.Logger.Debug("Service: OK")
	return &pb.CheckPermissionResponse{
		Code: int64(ErrCode.OK),
		Msg:  "ok",
		Res:  msg,
	}, nil
}
