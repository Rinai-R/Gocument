package main

import (
	"github.com/Rinai-R/Gocument/Server/Api/Initialize"
	"github.com/Rinai-R/Gocument/Server/Api/Router"
)

func main() {
	//开启路由
	Initialize.InitETCD()
	Router.InitRouter()
}
