package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/Api/Func/User/Client"
	"github.com/Rinai-R/Gocument/Server/Api/Func/User/Client/rpc"
	"github.com/Rinai-R/Gocument/pkg/Error/ErrCode"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/Rinai-R/Gocument/pkg/Rsp"
	"github.com/Rinai-R/Gocument/pkg/models"
	"github.com/Rinai-R/Gocument/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func PersonalPage(c context.Context, ctx *app.RequestContext) {
	_, exist := ctx.Get("GetName")
	if !exist {
		Logger.Logger.Debug("GetName is null")
		ctx.JSON(http.StatusUnauthorized, Rsp.TokenError("GetName not exist"))
		return
	}
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		Logger.Logger.Debug("BindError " + err.Error())
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err.Error()))
		return
	}
	res, _ := Client.UserClient.PersonalPage(c, &pb.PersonalPageRequest{
		UserId: user.Id,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		UserPage := utils.TimestampToTime(res)
		ctx.JSON(http.StatusOK, Rsp.Success(UserPage))
		break
	default:
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(res.Msg))
		break
	}
	return
}
