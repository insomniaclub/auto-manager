package router

import (
	v1 "auto-manager/api/v1"
	"auto-manager/middleware"
	"github.com/gin-gonic/gin"
)

func InitManagerRouter(r *gin.RouterGroup)  {
	ManagerRouter := r.Group("manager").Use(middleware.OperationRecord())
	{
		ManagerRouter.POST("register",v1.AddManager)
		ManagerRouter.POST("changePasswd",v1.ChangePasswd)
		ManagerRouter.DELETE("delete",v1.DelManager)
	}
}