package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/App/Func/Document/Client"
	pb "github.com/Rinai-R/Gocument/App/Func/Document/Client/rpc"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"log"
	"net/http"
)

func Enter(c context.Context, ctx *app.RequestContext) {

	GetName, exist := ctx.Get("GetName")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, Rsp.TokenError(nil))
	}
	var permission models.Permission
	err := ctx.BindJSON(&permission)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
		return
	}

	res, _ := Client.DocumentClient.Check(c, &pb.CheckPermissionRequest{
		Username:   GetName.(string),
		DocumentId: int64(permission.DocumentId),
	})
	code := int(res.Code)
	if code != ErrCode.OK {
		if code == ErrCode.DocumentNotFound {
			ctx.JSON(http.StatusBadRequest, Rsp.DocumentNotFound(nil))
			return
		}
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(res.Msg))
		return
	}
	if res.Res == "No Permission" {
		ctx.JSON(http.StatusForbidden, Rsp.EnterForbidden(nil))
		return
	}

	var upgrader = websocket.HertzUpgrader{} // 使用默认选项
	err = upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				// 处理读取消息时的错误
				ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(err.Error()))
				conn.Close()
				return
			}
			// 处理接收到的消息
			log.Printf("recv: %s", message)

			// 发送消息回客户端
			err = conn.WriteMessage(mt, message)
			if err != nil {
				// 处理发送消息时的错误
				ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(err.Error()))
				conn.Close()
				return
			}
		}
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(err.Error()))
		return
	}
}
