package main

import (
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/app/Router"
)

func main() {
	//开启日志
	Logger.InitLogger()
	//开启路由
	Router.InitRouter()
}
