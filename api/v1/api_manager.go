package v1

import (
	"auto-manager/model"
	"auto-manager/model/request"
	"auto-manager/model/response"
	"auto-manager/service"
	"github.com/gin-gonic/gin"
)

// @Tags Base
// @Summary 管理员登录
// @Produce application/json
// @Param data body request.Login true "nickname, password"
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"登录成功"}"
// @Router /v1/base/login [post]
func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	M := &model.Manager{
		Nickname: L.Nickname,
		Password: L.Password,
	}
	if err, _ := service.Login(M); err != nil {
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	response.OkWithMessage("登录成功", c)
}

func AddManager(c *gin.Context) {

}

func DelManager(c *gin.Context) {

}

func ChangePasswd(c *gin.Context) {

}

func GetManagerList(c *gin.Context) {

}

func token(c *gin.Context, m *model.Manager) {

}
