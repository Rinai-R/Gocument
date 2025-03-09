package Router

import (
	Api2 "github.com/Rinai-R/Gocument/Server/Api/Func/Document/Api"
	"github.com/Rinai-R/Gocument/Server/Api/Func/User/Api"
	"github.com/Rinai-R/Gocument/Server/Api/Middleware"
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

		UserGroup.GET("/page", Api.PersonalPage)
	}

	DocumentGroup := r.Group("/document")
	{
		DocumentGroup.Use(Middleware.Token())

		DocumentGroup.POST("/create", Api2.CreateDocument)

		DocumentGroup.PUT("/grant", Api2.Grant)

		DocumentGroup.DELETE("/delete", Api2.DeleteDocument)

		DocumentGroup.GET("/search", Api2.Search)

		DocumentGroup.GET("/enter", Api2.Enter)
	}

	r.Spin()
}
