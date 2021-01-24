package initializer

import (
	"auto-manager/global"
	"fmt"
	"github.com/fvbock/endless"
	"go.uber.org/zap"
	"time"
)

func RunServer() {
	r := InitRouters()
	port := fmt.Sprintf(":%d", global.AM_CONFIG.System.Port)
	s := endless.NewServer(port, r)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	time.Sleep(10 * time.Microsecond)
	global.AM_LOG.Info("Server run success on ", zap.String("port", port))
	fmt.Printf(`
	欢迎使用 Auto-Manager
	当前版本： developing
	默认自动化文档地址： http://127.0.0.1%s/swagger/index.html
	默认API运行地址： http://127.0.0.1%s/v1

`, port, port)

	global.AM_LOG.Error(s.ListenAndServe().Error())
}
