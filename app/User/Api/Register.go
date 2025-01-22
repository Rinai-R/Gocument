package Api

import (
	"context"
	"fmt"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/Utils/Error/ErrCode"
	"github.com/Rinai-R/Gocument/Utils/Rsp"
	"github.com/Rinai-R/Gocument/app/User/Client"
	pb "github.com/Rinai-R/Gocument/app/User/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Register(c context.Context, ctx *app.RequestContext) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Rsp.BindErr(err))
		return
	}
	res, _ := Client.UserClient.Register(c, &pb.RegisterRequest{
		UserName: user.Username,
		Password: user.Password,
	})
	fmt.Println(res)
	switch int(res.Code) {
	case ErrCode.OK:
		ctx.JSON(http.StatusOK, Rsp.Success(nil))
		break
	case ErrCode.UserNameLengthError:
		ctx.JSON(http.StatusBadRequest, Rsp.UserNameLengthErr(Error.ErrName))
		break
	case ErrCode.PasswordLengthError:
		ctx.JSON(http.StatusBadRequest, Rsp.PasswordLengthErr(Error.ErrpasswordLen))
		break
	case ErrCode.UserNameExists:
		ctx.JSON(http.StatusBadRequest, Rsp.UserNameExistsErr(Error.UserExists))
		break
	default:
		ctx.JSON(http.StatusInternalServerError, Rsp.InternalError(Error.InternalError))
		break
	}
	return
}
