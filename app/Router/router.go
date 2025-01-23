package Router

import (
	DApi "github.com/Rinai-R/Gocument/app/Func/Document/Api"
	"github.com/Rinai-R/Gocument/app/Func/User/Api"
	"github.com/Rinai-R/Gocument/app/Middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	r := server.Default()

	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/register", Api.Register)

		UserGroup.POST("/login", Api.Login)

		UserGroup.Use(Middleware.Token())

		UserGroup.PUT("/alter", Api.AlterUserInfo)
	}

	DocumentGroup := r.Group("/document")
	{
		DocumentGroup.Use(Middleware.Token())

		DocumentGroup.POST("/create", DApi.CreateDocument)

	}

	r.Spin()
}
