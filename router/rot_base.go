package router

import (
	v1 "auto-manager/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(r *gin.RouterGroup) {
	BaseRouter := r.Group("/base")
	{
		BaseRouter.POST("/token", v1.Login)
	}
}
