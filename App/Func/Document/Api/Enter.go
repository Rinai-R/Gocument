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
	"sync"
	"time"
)

var mu sync.Mutex

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
		Logger.Logger.Debug("Api: Document Not Found")
		msg, _ := json.Marshal(Rsp.DocumentNotFound(nil))
		_ = conn.WriteMessage(websocket.TextMessage, msg)
		break
	case ErrCode.OK:
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Only Read Mode"))
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
		//保存当前连接
		if models.Conns[DocumentId] == nil {
			Logger.Logger.Debug("Api: not Connected")
			models.Conns[DocumentId] = make(map[*websocket.Conn]bool)
		}
		models.Conns[DocumentId][conn] = true // 将连接添加到相应文档组
		//退出之前连接取消
		defer func() {
			delete(models.Conns[DocumentId], conn)
			Logger.Logger.Debug("Api: Connection Disconnected")
		}()
		select {}
	default:
		Logger.Logger.Debug("Api: Internal Error " + res.Msg)
		msg, _ := json.Marshal(Rsp.InternalError(res.Msg))
		_ = conn.WriteMessage(websocket.TextMessage, msg)
		break
	}
	return
}

// ReadAndWrite 读写操作，基本逻辑和只读(Only Read)一致，但多了一个读取前端的输入框的处理。
func ReadAndWrite(c context.Context, _ *app.RequestContext, DocumentId int, conn *websocket.Conn) {
	res, _ := Client.DocumentClient.Get(c, &pb.GetDocumentRequest{
		DocumentId: int64(DocumentId),
	})
	switch int(res.Code) {
	case ErrCode.DocumentNotFound:
		Logger.Logger.Debug("Api: Document Not Found")
		msg, _ := json.Marshal(Rsp.DocumentNotFound(nil))
		_ = conn.WriteMessage(websocket.TextMessage, msg)
		break

		//OK的逻辑
	case ErrCode.OK:
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Read And Write Mode"))
		//先将结构体转成json格式发送给前端
		Document := &models.ElasticDocument{
			Id:       strconv.Itoa(DocumentId),
			Title:    res.Title,
			Content:  res.Content,
			CreateAt: res.CreateAt.AsTime(),
			UpdateAt: res.UpdatedAt.AsTime(),
		}
		mes, _ := json.Marshal(Document)
		err := conn.WriteMessage(websocket.TextMessage, mes)
		if err != nil {
			msg, _ := json.Marshal(Rsp.InternalError(err.Error()))
			_ = conn.WriteMessage(websocket.TextMessage, msg)
			return
		}
		//保存当前连接
		if models.Conns[DocumentId] == nil {
			Logger.Logger.Debug("Api: not Connected")
			models.Conns[DocumentId] = make(map[*websocket.Conn]bool)
		}
		models.Conns[DocumentId][conn] = true // 将连接添加到相应文档组
		//退出之前连接取消
		defer func() {
			delete(models.Conns[DocumentId], conn)
			Logger.Logger.Debug("Api: Connection Disconnected")
		}()

		for {
			//保存逻辑
			//先读取信息
			_, msg, err := conn.ReadMessage()
			if err != nil {
				//错误处理
				Logger.Logger.Debug("Api: Internal Error " + err.Error())
				msg, _ := json.Marshal(Rsp.InternalError(err.Error()))
				_ = conn.WriteMessage(websocket.TextMessage, msg)
				return
			}

			//此处将要向服务端发送保存的信息
			var document models.ElasticDocument
			//将读取的信息绑定在结构体上
			err = json.Unmarshal(msg, &document)
			if err != nil {
				msg, _ := json.Marshal(Rsp.InternalError(err.Error()))
				_ = conn.WriteMessage(websocket.TextMessage, msg)
			}
			document.UpdateAt = time.Now()
			//发送请求
			res, _ := Client.DocumentClient.Edit(c, &pb.EditRequest{
				DocumentId: int64(DocumentId),
				Title:      document.Title,
				Content:    document.Content,
			})
			//根据错误码判断是否成功
			switch int(res.Code) {
			case ErrCode.DocumentNotFound:
				mes, _ := json.Marshal(Rsp.DocumentNotFound(nil))
				_ = conn.WriteMessage(websocket.TextMessage, mes)
				break
			case ErrCode.OK:
				Logger.Logger.Debug("Api: OK")
				//ok吗噜了 然后广播这个保存的信息
				mes, _ := json.Marshal(document)
				broadcast(DocumentId, mes)
				break
			case ErrCode.SensitiveWords:
				mes, _ := json.Marshal(Rsp.SensitiveWords(nil))
				_ = conn.WriteMessage(websocket.TextMessage, mes)
				break

			}
		}

	default:
		Logger.Logger.Debug("Api: Internal Error " + res.Msg)
		msg, _ := json.Marshal(Rsp.InternalError(res.Msg))
		_ = conn.WriteMessage(websocket.TextMessage, msg)
		break
	}
	return
}

func broadcast(DocumentId int, mes []byte) {
	clients, exists := models.Conns[DocumentId]
	if !exists {
		return
	}
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, mes)
		if err != nil {
			Logger.Logger.Debug("Api: Broadcast Error " + err.Error())
			client.Close()          // 关闭无法写入的连接
			delete(clients, client) // 从列表中删除
		}
	}
}
