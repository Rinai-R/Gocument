package handle

import (
	"context"
	"errors"
	dao2 "github.com/Rinai-R/Gocument/Server/Document/DataBase/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/pkg/Error"
	"github.com/Rinai-R/Gocument/pkg/Error/ErrCode"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/Rinai-R/Gocument/pkg/models"
)

func (*DocumentServer) Delete(c context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, err := dao2.GetId(c, req.Username)
	if err != nil {
		Logger.Logger.Debug("Service: Unknown username " + err.Error())
	}
	document := models.Document{
		Id:     int(req.Id),
		UserId: id,
	}

	if err = dao2.Delete(c, document); err != nil {
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
