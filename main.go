package main

import (
	"auto-manager/global"
	"auto-manager/initializer"
	"fmt"
)

func main()  {
	fmt.Println("Hello")
	global.AM_VP = initializer.InitViper()
	global.AM_LOG = initializer.InitZap()
	global.AM_DB = initializer.InitGorm()

	db,_ := global.AM_DB.DB()
	defer db.Close()

	initializer.RunServer()
}
