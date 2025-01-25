package service

import (
	"context"
	"github.com/Rinai-R/Gocument/DataBase/Document/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/models"
	"strconv"
)

func (*DocumentServer) Edit(ctx context.Context, req *pb.EditRequest) (*pb.EditResponse, error) {
	document := models.ElasticDocument{
		Id:      strconv.FormatInt(req.DocumentId, 10),
		Title:   req.Title,
		Content: req.Content,
	}
	if err := dao.Edit(ctx, document); err != nil {
		return &pb.EditResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}
	return &pb.EditResponse{
		Code: int64(ErrCode.OK),
		Msg:  "ok",
	}, nil
}
