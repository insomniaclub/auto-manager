package v1

import (
	"auto-manager/model"
	"auto-manager/model/request"
	"auto-manager/model/response"
	"auto-manager/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var L request.LoginRequest
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

func token(m *model.Manager) {

}
