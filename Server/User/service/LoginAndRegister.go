package service

import (
	"context"
	"errors"
	"github.com/Rinai-R/Gocument/DataBase/User/dao"
	"github.com/Rinai-R/Gocument/Logger"
	pb "github.com/Rinai-R/Gocument/Server/User/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/encrypt"
	"github.com/Rinai-R/Gocument/models"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

// Register 重写注册方法
func (*UserServer) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := models.User{
		Username: req.UserName,
		Password: req.Password,
	}

	if len(user.Username) > 20 || len(user.Username) < 5 {
		Logger.Logger.Debug("Service: User Username Length Error")
		return &pb.RegisterResponse{
			Code: int64(ErrCode.UserNameLengthError),
			Msg:  Error.ErrName.Error(),
		}, nil
	}

	if len(user.Password) < 5 || len(user.Password) > 25 {
		Logger.Logger.Debug("Service: User Password Length Error")
		return &pb.RegisterResponse{
			Code: int64(ErrCode.PasswordLengthError),
			Msg:  Error.ErrpasswordLen.Error(),
		}, nil
	}
	//加密存储
	user.Password = encrypt.EncryptPassword(user.Password)

	//注册到数据库时出现错误
	if err := dao.Register(user); err != nil {
		if errors.Is(err, Error.UserExists) {
			Logger.Logger.Debug("Service: User Exists Error")
			return &pb.RegisterResponse{
				Code: int64(ErrCode.UserNameExists),
				Msg:  Error.UserExists.Error(),
			}, nil
		}
		Logger.Logger.Debug("Service: Internal Error" + err.Error())
		return &pb.RegisterResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  Error.InternalError.Error(),
		}, nil
	}
	Logger.Logger.Debug("Service: User Success")
	return &pb.RegisterResponse{
		Code: int64(ErrCode.OK),
		Msg:  "OK",
	}, nil
}

func (*UserServer) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := models.User{
		Username: req.UserName,
		Password: req.Password,
	}

	if err := dao.Login(user); err != nil {
		if errors.Is(err, Error.UsernameOrPassword) {
			Logger.Logger.Debug("Service: User Username Or Password Error")
			return &pb.LoginResponse{
				Code: int64(ErrCode.UsernameOrPassword),
				Msg:  Error.UsernameOrPassword.Error(),
			}, nil
		}
		Logger.Logger.Debug("Service: Internal Error" + err.Error())
		return &pb.LoginResponse{
			Code: int64(ErrCode.InternalErr),
			Msg:  err.Error(),
		}, nil
	}
	Logger.Logger.Debug("Service: User Success")
	return &pb.LoginResponse{
		Code: int64(ErrCode.OK),
		Msg:  "Msg",
	}, nil

}
