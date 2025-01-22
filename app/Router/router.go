package Router

import (
	"github.com/Rinai-R/Gocument/app/User/Api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	r := server.Default()

	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/register", Api.Register)

		UserGroup.POST("/login", Api.Login)
	}

	r.Spin()
}
