package Api

import (
	"context"
	"encoding/json"
	"github.com/Rinai-R/Gocument/App/Func/Document/Client"
	pb "github.com/Rinai-R/Gocument/App/Func/Document/Client/rpc"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"net/http"
	"strconv"
	"time"
)

// Enter websocket请求，进入文档界面，包括禁止进入，只读，读写的权限处理
func Enter(c context.Context, ctx *app.RequestContext) {
	var upgrader = websocket.HertzUpgrader{
		CheckOrigin: func(ctx *app.RequestContext) bool {
			return true
		},
	}
	err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		//先处理信息，再进行逻辑判断
		GetName, exist := ctx.Get("GetName")
		if !exist {
			Logger.Logger.Debug("Api: GetName not exist")
			msg, _ := json.Marshal(Rsp.TokenError("GetName not exist"))
			_ = conn.WriteMessage(websocket.TextMessage, msg)
			return
		}
		//拿到访问文档的id信息
		documentID := ctx.Query("document_id")
		var permission models.Permission
		var err error
		permission.DocumentId, err = strconv.Atoi(documentID)
		if documentID == "" || err != nil {
			Logger.Logger.Debug("Api: documentID Bind Error")
			msg, _ := json.Marshal(Rsp.BindErr("DocumentID BindErr"))
			_ = conn.WriteMessage(websocket.TextMessage, msg)
			return
		}
		//拿到用户相关的权限信息
		res, _ := Client.DocumentClient.Check(c, &pb.CheckPermissionRequest{
			Username:   GetName.(string),
			DocumentId: int64(permission.DocumentId),
		})
		//处理错误
		code := int(res.Code)
		if code != ErrCode.OK {
			if code == ErrCode.DocumentNotFound {
				Logger.Logger.Debug("Api: Document Not Found")
				msg, _ := json.Marshal(Rsp.DocumentNotFound(nil))
				_ = conn.WriteMessage(websocket.TextMessage, msg)
				return
			}
			Logger.Logger.Debug("Api: Internal Error " + res.Msg)
			msg, _ := json.Marshal(Rsp.InternalError(res.Msg))
			_ = conn.WriteMessage(websocket.TextMessage, msg)
			return
		}

		//根据得到的权限类型走分支：
		switch res.Res {
		case "No Permission":
			Logger.Logger.Debug("Api: No Permission")
			msg, _ := json.Marshal(Rsp.EnterForbidden(nil))
			_ = conn.WriteMessage(websocket.TextMessage, msg)
			break
		case "Only Read":
			Logger.Logger.Debug("Api: Only Read")
			OnlyRead(c, ctx, permission.DocumentId, conn)
			break
		case "RW":
			Logger.Logger.Debug("Api: Read And Write")
			ReadAndWrite(c, ctx, permission.DocumentId, conn)
			break
		}
		return
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(err.Error()))
		return
	}
}

// OnlyRead 只读状态下的进入文档
func OnlyRead(c context.Context, ctx *app.RequestContext, DocumentId int, conn *websocket.Conn) {
	res, _ := Client.DocumentClient.Get(c, &pb.GetDocumentRequest{
		DocumentId: int64(DocumentId),
	})
	switch int(res.Code) {
	case ErrCode.DocumentNotFound:
		msg, _ := json.Marshal(Rsp.DocumentNotFound(nil))
		_ = conn.WriteMessage(websocket.TextMessage, msg)
		break
	case ErrCode.OK:
		Document := &models.ElasticDocument{
			Id:      strconv.Itoa(DocumentId),
			Title:   res.Title,
			Content: res.Content,
		}
		mes, _ := json.Marshal(Document)
		err := conn.WriteMessage(websocket.TextMessage, mes)
		if err != nil {
			msg, _ := json.Marshal(Rsp.InternalError(err.Error()))
			_ = conn.WriteMessage(websocket.TextMessage, msg)
			return
		}
		time.Sleep(5 * time.Second)
		defer OnlyRead(c, ctx, DocumentId, conn)
	default:
		Logger.Logger.Debug("Api: Internal Error " + res.Msg)
		msg, _ := json.Marshal(Rsp.InternalError(res.Msg))
		_ = conn.WriteMessage(websocket.TextMessage, msg)
		break
	}
	return
}

func ReadAndWrite(c context.Context, ctx *app.RequestContext, DocumentId int, conn *websocket.Conn) {

}
