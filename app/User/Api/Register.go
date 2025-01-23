package Api

import (
	"context"
	"fmt"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/app/User/Client"
	pb "github.com/Rinai-R/Gocument/app/User/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Register(c context.Context, ctx *app.RequestContext) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		Logger.Logger.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
		return
	}
	res, _ := Client.UserClient.Register(c, &pb.RegisterRequest{
		UserName: user.Username,
		Password: user.Password,
	})
	fmt.Println(res)
	switch int(res.Code) {
	case ErrCode.OK:
		Logger.Logger.Info("API: register success")
		ctx.JSON(http.StatusOK, Rsp.Success(nil))
		break
	case ErrCode.UserNameLengthError:
		Logger.Logger.Info("API: user name length error")
		ctx.JSON(http.StatusBadRequest, Rsp.UserNameLengthErr(Error.ErrName))
		break
	case ErrCode.PasswordLengthError:
		Logger.Logger.Info("API: password length error")
		ctx.JSON(http.StatusBadRequest, Rsp.PasswordLengthErr(Error.ErrpasswordLen))
		break
	case ErrCode.UserNameExists:
		Logger.Logger.Info("API: user name exists")
		ctx.JSON(http.StatusBadRequest, Rsp.UserNameExistsErr(Error.UserExists))
		break
	default:
		Logger.Logger.Error("API: InternalError")
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(Error.InternalError))
		break
	}
	return
}
