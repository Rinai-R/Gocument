package service

import (
	"context"
	"errors"
	"github.com/Rinai-R/Gocument/DataBase/Document/dao"
	"github.com/Rinai-R/Gocument/Logger"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/models"
)

func (*DocumentServer) Delete(c context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, err := dao.GetId(req.Username)
	if err != nil {
		Logger.Logger.Debug("Service: Unknown username " + err.Error())
	}
	document := models.Document{
		Id:     int(req.Id),
		UserId: id,
	}

	if err = dao.Delete(c, document); err != nil {
		Logger.Logger.Debug("Service: Delete Error " + err.Error())
		if errors.Is(err, Error.NoDocumentFoundWithToken) {
			return &pb.DeleteResponse{
				Code: int64(ErrCode.NoDocumentFoundWithToken),
				Msg:  Error.NoDocumentFoundWithToken.Error(),
			}, nil
		}
		return &pb.DeleteResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}
	return &pb.DeleteResponse{
		Code: int64(ErrCode.OK),
		Msg:  "ok",
	}, nil
}
