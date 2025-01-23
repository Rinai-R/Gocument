package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/Utils/utils"
	"github.com/Rinai-R/Gocument/app/Func/User/Client"
	pb "github.com/Rinai-R/Gocument/app/Func/User/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func PersonalPage(c context.Context, ctx *app.RequestContext) {
	GetName, exist := ctx.Get("GetName")
	if !exist {
		Logger.Logger.Debug("GetName is null")
		ctx.JSON(http.StatusUnauthorized, Rsp.TokenError("GetName not exist"))
		return
	}
	var user models.User
	user.Username = GetName.(string)
	res, _ := Client.UserClient.PersonalPage(c, &pb.PersonalPageRequest{
		Username: user.Username,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		if UserPage, ok := utils.TimestampToTime(res); ok {
			ctx.JSON(http.StatusOK, Rsp.Success(UserPage))
		} else {
			ctx.JSON(http.StatusInternalServerError, Rsp.InternalError("TimeTransition failed"))
		}
		break
	default:
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(res.Msg))
		break
	}
	return
}
