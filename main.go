package main

import (
	"auto-manager/global"
	"auto-manager/initializer"
	"fmt"
)

// @title Auto-Manager Dev API
// @version 0.0.1
// @description This is swagger sample for auto-manager
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
