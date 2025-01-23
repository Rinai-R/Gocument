package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/app/Middleware"
	"github.com/Rinai-R/Gocument/app/User/Client"
	pb "github.com/Rinai-R/Gocument/app/User/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Login(c context.Context, ctx *app.RequestContext) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		Logger.Logger.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
	}
	res, _ := Client.UserClient.Login(c, &pb.LoginRequest{
		UserName: user.Username,
		Password: user.Password,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		Logger.Logger.Debug("API: login ok")
		token, err := Middleware.GenerateJWT(user.Username)
		if err != nil {
			Logger.Logger.Debug("API: " + err.Error())
			token = "token generate failed"
		}
		ctx.JSON(http.StatusOK, Rsp.Success(token))
		break
	case ErrCode.UsernameOrPassword:
		Logger.Logger.Debug("API: login username or password error")
		ctx.JSON(http.StatusBadRequest, Rsp.UsernameOrPassword(ErrCode.UsernameOrPassword))
		break
	default:
		Logger.Logger.Debug("API: Internal error" + res.Msg)
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(res.Msg))
		break
	}
	return
}
