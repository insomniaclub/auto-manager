package router

import (
	v1 "auto-manager/api/v1"
	"github.com/gin-gonic/gin"
)

func InitManagerRouter(r *gin.RouterGroup) {
	// TODO 操作记录
	ManagerRouter := r.Group("/manager") //.Use(middleware.OperationRecord())
	{
		ManagerRouter.POST("/", v1.AddManager)
		ManagerRouter.PUT("/:id/passwd", v1.ChangePasswd)
		ManagerRouter.DELETE("/:id", v1.DeleteManager)
		ManagerRouter.GET("/", v1.GetManagerList)
	}
}
