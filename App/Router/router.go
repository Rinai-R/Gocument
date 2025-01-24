package Router

import (
	DApi "github.com/Rinai-R/Gocument/App/Func/Document/Api"
	UApi "github.com/Rinai-R/Gocument/App/Func/User/Api"
	"github.com/Rinai-R/Gocument/App/Middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() {
	r := server.Default()

	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/register", UApi.Register)

		UserGroup.POST("/login", UApi.Login)

		UserGroup.Use(Middleware.Token())

		UserGroup.PUT("/alter", UApi.AlterUserInfo)

		UserGroup.GET("/page", UApi.PersonalPage)
	}

	DocumentGroup := r.Group("/document")
	{
		DocumentGroup.Use(Middleware.Token())

		DocumentGroup.POST("/create", DApi.CreateDocument)

		DocumentGroup.DELETE("/delete", DApi.DeleteDocument)

		DocumentGroup.PUT("/enter", DApi.Enter)
	}

	r.Spin()
}
