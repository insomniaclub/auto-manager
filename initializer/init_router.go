package initializer

import (
	_ "auto-manager/docs"
	"auto-manager/middleware"
	"auto-manager/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouters() (r *gin.Engine) {
	r = gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	PublicGroup := r.Group("v1")
	{
		router.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := r.Group("v1")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.InterceptorHandle())
	{
		router.InitManagerRouter(PrivateGroup)
	}
	return r
}
