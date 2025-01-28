package service

import (
	"context"
	dao2 "github.com/Rinai-R/Gocument/Server/Document/DataBase/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error"
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
	if !dao2.SensitiveCheck(ctx, req.Title) || !dao2.SensitiveCheck(ctx, req.Content) {
		return &pb.EditResponse{
			Code: int64(ErrCode.SensitiveWords),
			Msg:  Error.SensitiveWords.Error(),
		}, nil
	}
	if err := dao2.Edit(ctx, document); err != nil {
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
