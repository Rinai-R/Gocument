package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/App/Func/Document/Client"
	pb "github.com/Rinai-R/Gocument/App/Func/Document/Client/rpc"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/Utils/utils"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Search(c context.Context, ctx *app.RequestContext) {
	_, exist := ctx.Get("GetName")
	if !exist {
		Logger.Logger.Debug("GetName not exist")
		ctx.JSON(http.StatusBadRequest, Rsp.TokenError("GetName not exist"))
		return
	}

	var SearchRequest models.SearchDocument
	err := ctx.BindJSON(&SearchRequest)
	if err != nil {
		Logger.Logger.Debug("Bind Error " + err.Error())
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr("Bind Error "+err.Error()))
		return
	}
	res, _ := Client.DocumentClient.Search(c, &pb.SearchRequest{
		Content: SearchRequest.Content,
		UserId:  SearchRequest.UserId,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		Logger.Logger.Debug("Api: Search Success")
		ans := utils.TimeTransition(res)
		ctx.JSON(http.StatusOK, Rsp.Success(ans))
		break
	case ErrCode.SearchError:
		Logger.Logger.Error("Api: Search Error")
		ctx.JSON(http.StatusBadRequest, Rsp.SearchError("Search Error"))
		break
	}
	return

}
