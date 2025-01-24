package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/App/Func/User/Client"
	"github.com/Rinai-R/Gocument/App/Func/User/Client/rpc"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func AlterUserInfo(c context.Context, ctx *app.RequestContext) {
	var user models.User
	GetName, exist := ctx.Get("GetName")
	if !exist {
		Logger.Logger.Debug("Api: Token null")
		ctx.JSON(http.StatusUnauthorized, Rsp.TokenError(nil))
		return
	}
	err := ctx.BindJSON(&user)
	if err != nil {
		Logger.Logger.Debug("Api: BindJSON Error " + err.Error())
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
		return
	}
	user.Username = GetName.(string)
	res, _ := Client.UserClient.Alter(c, &pb.AlterRequest{
		Username: user.Username,
		Password: user.Password,
		Bio:      user.Bio,
		Gender:   user.Gender,
		Avatar:   user.Avatar,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		Logger.Logger.Debug("API: AlterUserInfo ok")
		ctx.JSON(http.StatusOK, Rsp.Success(nil))
		break
	case ErrCode.RequestNull:
		Logger.Logger.Debug("API: AlterUserInfo Request null")
		ctx.JSON(http.StatusBadRequest, Rsp.RequestNull(nil))
		break
	case ErrCode.UserNotExists:
		Logger.Logger.Debug("API: AlterUserInfo UserNotExists")
		ctx.JSON(http.StatusBadRequest, Rsp.UserNotExists(nil))
		break
	default:
		Logger.Logger.Debug("API: Internal Error " + res.Msg)
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(res.Msg))
		break
	}
	return
}
