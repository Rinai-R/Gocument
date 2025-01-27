package service

import (
	"context"
	"errors"
	Ddao "github.com/Rinai-R/Gocument/DataBase/Document/dao"
	Udao "github.com/Rinai-R/Gocument/DataBase/User/dao"
	"github.com/Rinai-R/Gocument/Logger"
	pb "github.com/Rinai-R/Gocument/Server/User/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/encrypt"
	"github.com/Rinai-R/Gocument/models"
)

func (*UserServer) Alter(ctx context.Context, req *pb.AlterRequest) (*pb.AlterResponse, error) {
	if req.Avatar == "" && req.Bio == "" && req.Password == "" && req.Gender == "" {
		return &pb.AlterResponse{
			Code: int64(ErrCode.RequestNull),
			Msg:  Error.RequestNull.Error(),
		}, nil
	}
	if !Ddao.SensitiveCheck(ctx, req.Bio) || !Ddao.SensitiveCheck(ctx, req.Gender) {
		return &pb.AlterResponse{
			Code: int64(ErrCode.SensitiveWords),
			Msg:  Error.SensitiveWords.Error(),
		}, nil
	}
	if req.Password != "" && len(req.Password) < 5 || len(req.Password) > 25 {
		return &pb.AlterResponse{
			Code: int64(ErrCode.PasswordLengthError),
			Msg:  Error.ErrpasswordLen.Error(),
		}, nil
	}

	req.Password = encrypt.EncryptPassword(req.Password)

	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Bio:      req.Bio,
		Gender:   req.Gender,
		Avatar:   req.Avatar,
	}

	if err := Udao.AlterUserInfo(user); err != nil {
		if errors.Is(err, Error.UserNotExists) {
			Logger.Logger.Debug("Service: Not Find User")
			return &pb.AlterResponse{
				Code: int64(ErrCode.UserNotExists),
				Msg:  Error.UserNotExists.Error(),
			}, nil
		}
		Logger.Logger.Debug("Service: Internal Error, Alter User Failed " + err.Error())
		return &pb.AlterResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  Error.InternalError.Error(),
		}, nil
	}
	return &pb.AlterResponse{
		Code: int64(ErrCode.OK),
		Msg:  "OK",
	}, nil
}
