package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/app/Func/Document/Client"
	pb "github.com/Rinai-R/Gocument/app/Func/Document/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func DeleteDocument(c context.Context, ctx *app.RequestContext) {
	GetName, exist := ctx.Get("GetName")
	if !exist {
		Logger.Logger.Debug("Api: Token NULL")
		ctx.JSON(http.StatusUnauthorized, Rsp.TokenError(nil))
		return
	}
	var document models.Document

	err := ctx.BindJSON(&document)
	if err != nil {
		Logger.Logger.Debug("Api: bind error")
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
		return
	}
	res, _ := Client.DocumentClient.Delete(c, &pb.DeleteRequest{
		Id:       int64(document.Id),
		Username: GetName.(string),
	})
	switch int(res.Code) {
	case ErrCode.OK:
		Logger.Logger.Debug("Api: Delete OK")
		ctx.JSON(http.StatusOK, Rsp.Success(nil))
		break
	case ErrCode.NoDocumentFoundWithToken:
		Logger.Logger.Debug("Api: " + Error.NoDocumentFoundWithToken.Error())
		ctx.JSON(http.StatusBadRequest, Rsp.NoDocumentFoundWithToken(nil))
		break
	default:
		Logger.Logger.Debug("Api: Delete Error Internal " + res.Msg)
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(res.Msg))
		break
	}
	return

}
