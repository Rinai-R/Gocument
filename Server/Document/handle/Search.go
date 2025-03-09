package handle

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/dao"
	pb "github.com/Rinai-R/Gocument/Server/Document/rpc"
	"github.com/Rinai-R/Gocument/pkg/Error"
	"github.com/Rinai-R/Gocument/pkg/Error/ErrCode"
	"github.com/Rinai-R/Gocument/pkg/models"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

func (*DocumentServer) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	Doc := models.SearchDocument{
		UserId:  req.UserId,
		Content: req.Content,
	}
	if res := dao.Search(ctx, Doc); res != nil {
		var doc []*pb.Doc
		for _, slice := range res {
			id, _ := strconv.Atoi(slice.Id)
			doc = append(doc, &pb.Doc{
				Id:        int64(id),
				Title:     slice.Title,
				IsPrivate: slice.IsPrivate,
				CreateAt:  timestamppb.New(slice.CreateAt),
				UpdatedAt: timestamppb.New(slice.UpdateAt),
			})
		}

		return &pb.SearchResponse{
			Code:      int64(ErrCode.OK),
			Msg:       "ok",
			Documents: doc,
		}, nil
	}
	return &pb.SearchResponse{
		Code: int64(ErrCode.SearchError),
		Msg:  Error.SearchError.Error(),
	}, nil

}
