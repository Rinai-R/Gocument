package handle

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/User/DataBase/dao"
	pb "github.com/Rinai-R/Gocument/Server/User/rpc"
	"github.com/Rinai-R/Gocument/pkg/Error/ErrCode"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/Rinai-R/Gocument/pkg/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (*UserServer) PersonalPage(c context.Context, req *pb.PersonalPageRequest) (*pb.PersonalPageResponse, error) {
	user := models.User{
		Id: req.UserId,
	}
	if err := dao.PersonalPage(c, &user); err != nil {
		return &pb.PersonalPageResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
			User: nil,
		}, nil
	}
	res := &pb.UserInfo{
		Id:        user.Id,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		Gender:    user.Gender,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	for _, d := range user.Documents {
		res.Documents = append(res.Documents, &pb.Do{
			Id:        int64(d.Id),
			Title:     d.Title,
			IsPrivate: d.IsPrivate,
			CreateAt:  timestamppb.New(d.CreateAt),
			UpdatedAt: timestamppb.New(d.UpdateAt),
		})
	}
	Logger.Logger.Debug("response generate ok")

	return &pb.PersonalPageResponse{
		Code: int64(ErrCode.OK),
		Msg:  "ok",
		User: res,
	}, nil
}
