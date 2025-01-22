package Api

import (
	"context"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/app/Middleware"
	"github.com/Rinai-R/Gocument/app/User/Client"
	pb "github.com/Rinai-R/Gocument/app/User/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Login(c context.Context, ctx *app.RequestContext) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
	}
	res, _ := Client.UserClient.Login(c, &pb.LoginRequest{
		UserName: user.Username,
		Password: user.Password,
	})
	switch int(res.Code) {
	case ErrCode.OK:
		token, err := Middleware.GenerateJWT(user.Username)
		if err != nil {
			token = "token generate failed"
		}
		ctx.JSON(http.StatusOK, Rsp.Success(token))
		break
	case ErrCode.UsernameOrPassword:
		ctx.JSON(http.StatusBadRequest, Rsp.UsernameOrPassword(ErrCode.UsernameOrPassword))
		break
	default:
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(ErrCode.InternalErr))
		break
	}
	return
}
