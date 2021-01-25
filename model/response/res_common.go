package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, "", msg, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, "", msg, c)
}
