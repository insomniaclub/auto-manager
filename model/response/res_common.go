package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, Response{
		Data: data,
		Msg:  msg,
	})
}

func WithMessage(code int, msg string, c *gin.Context) {
	Result(code, "", msg, c)
}

func WithDetail(code int, data interface{}, msg string, c *gin.Context) {
	Result(code, data, msg, c)
}
