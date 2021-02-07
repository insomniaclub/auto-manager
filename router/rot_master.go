package router

import (
	v1 "auto-manager/api/v1"
	"github.com/gin-gonic/gin"
)

func InitMasterRouter(r *gin.RouterGroup) {
	MasterRouter := r.Group("/master")
	{
		MasterRouter.POST("/token", v1.WxLogin)
	}
}
