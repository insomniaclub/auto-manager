package main

import (
	"auto-manager/global"
	"auto-manager/initializer"
	"fmt"
)

// @title Auto-Manager Dev API
// @version 0.0.1
// @description Auto-Manager 的接口设置完全遵循Restful原则，具体参考API文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	fmt.Println("Hello")
	global.AM_VP = initializer.InitViper()
	global.AM_LOG = initializer.InitZap()
	global.AM_DB = initializer.InitGorm()

	db, _ := global.AM_DB.DB()
	defer db.Close()

	initializer.RunServer()
}
