package service

import (
	"context"
	"errors"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/models"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

func (*DocumentServer) Get(ctx context.Context, req *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	ESDocument := models.ElasticDocument{
		Id: strconv.FormatInt(req.DocumentId, 10),
	}
	if err := dao.GetDocument(ctx, &ESDocument); err != nil {
		if errors.Is(err, Error.DocumentNotFound) {
			Logger.Logger.Debug("Service: Document not found")
			return &pb.GetDocumentResponse{
				Code: int64(ErrCode.DocumentNotFound),
				Msg:  Error.DocumentNotFound.Error(),
			}, nil
		}
		Logger.Logger.Debug("Service: Internal Error" + err.Error())
		return &pb.GetDocumentResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}
	Logger.Logger.Debug("Service: OK")
	return &pb.GetDocumentResponse{
		Code:      int64(ErrCode.OK),
		Msg:       "ok",
		Title:     ESDocument.Title,
		Content:   ESDocument.Content,
		CreateAt:  timestamppb.New(ESDocument.CreateAt),
		UpdatedAt: timestamppb.New(ESDocument.UpdateAt),
	}, nil
}
