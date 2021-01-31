package router

import (
	v1 "auto-manager/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPositionRouter(r *gin.RouterGroup) {
	PositionRouter := r.Group("/position")
	{
		PositionRouter.GET("/", v1.GetPositionList)
		PositionRouter.POST("/", v1.AddPosition)
		PositionRouter.DELETE("/:id", v1.DeletePosition)
		PositionRouter.PUT("/:id", v1.ModifyPosition)
	}
}
