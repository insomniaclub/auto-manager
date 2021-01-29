package v1

import (
	"auto-manager/global"
	"auto-manager/model"
	"auto-manager/model/request"
	"auto-manager/model/response"
	"auto-manager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 管理员登录
// @Produce application/json
// @Param data body request.Login true "nickname, password"
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"登录成功"}"
// @Router /v1/base/login [post]
func Login(c *gin.Context) {
	var reqL request.Login
	_ = c.ShouldBindJSON(&reqL)
	m := &model.Manager{
		Nickname: reqL.Nickname,
		Password: reqL.Password,
	}
	resL, err := service.Login(m)
	if err != nil {
		global.AM_LOG.Error("Login Failed.", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetail(resL, "登录成功", c)
}

func AddManager(c *gin.Context) {

}

func DelManager(c *gin.Context) {

}

func ChangePasswd(c *gin.Context) {

}

func GetManagerList(c *gin.Context) {

}
