package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/App/Func/Document/Client"
	pb "github.com/Rinai-R/Gocument/App/Func/Document/Client/rpc"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func CreateDocument(c context.Context, ctx *app.RequestContext) {
	GetName, exists := ctx.Get("GetName")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, Rsp.TokenError(nil))
		return
	}
	var Document models.Document
	err := ctx.BindJSON(&Document)
	if err != nil || Document.Title == "" {
		if Document.Title == "" {
			ctx.JSON(http.StatusUnauthorized, Rsp.BindErr("Title is required"))
			return
		}
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
		return
	}
	res, _ := Client.DocumentClient.Create(c, &pb.CreateRequest{
		Username:  GetName.(string),
		Title:     Document.Title,
		IsPrivate: Document.IsPrivate,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		Logger.Logger.Debug("Create Document OK")
		ctx.JSON(http.StatusOK, Rsp.Success(nil))
		break
	default:
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(res.Msg))
		break
	}
	return
}
