package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/App/Func/Document/Client"
	pb "github.com/Rinai-R/Gocument/App/Func/Document/Client/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// Grant 文档授权，如果请求的用户具有该权限，则为取消权限，如果没有，则添加该权限。
func Grant(c context.Context, ctx *app.RequestContext) {
	GetName, exist := ctx.Get("GetName")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, Rsp.TokenError(nil))
		return
	}
	var Permission models.Permission

	err := ctx.BindJSON(&Permission)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
		return
	}

	res, _ := Client.DocumentClient.Grant(c, &pb.GrantRequest{
		Host:       GetName.(string),
		UserId:     int64(Permission.UserId),
		DocumentId: int64(Permission.DocumentId),
		Type:       Permission.Type,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		ctx.JSON(http.StatusOK, Rsp.Success(nil))
		break
	default:
		ctx.JSON(http.StatusBadRequest, Rsp.GrantFailed(res.Msg))
		break
	}
	return
}
