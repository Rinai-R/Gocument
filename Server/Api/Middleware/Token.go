package Middleware

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/Api/Func/Auth/fn"
	"github.com/Rinai-R/Gocument/pkg/Rsp"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Token() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		auth := ctx.Request.Header.Get("Authorization")

		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, Rsp.TokenError("token null"))
			ctx.Abort()
			return
		}

		claims, err := fn.ParseToken(auth)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, Rsp.TokenError(err))
			ctx.Abort()
			return
		}
		ctx.Set("GetName", claims)

		ctx.Next(c)
	}
}
