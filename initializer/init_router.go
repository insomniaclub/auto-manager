package initializer

import (
	"auto-manager/middleware"
	"auto-manager/router"
	"github.com/gin-gonic/gin"
)

func InitRouters() (r *gin.Engine) {
	r = gin.Default()
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
