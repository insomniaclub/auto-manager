package v1

import (
	"auto-manager/global"
	"auto-manager/model"
	"auto-manager/model/request"
	"auto-manager/model/response"
	"auto-manager/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// @Tags base
// @Summary 管理员登录
// @Produce application/json
// @Param data body request.Login true "nickname, password"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"登录成功"}"
// @Router /v1/base/token [post]
func Login(c *gin.Context) {
	var req request.Login
	_ = c.ShouldBindJSON(&req)
	m := &model.Manager{
		Nickname: req.Nickname,
		Password: req.Password,
	}
	res, err := service.Login(m)
	if err != nil {
		global.AM_LOG.Error("Login Failed.", zap.Any("err", err))
		response.WithMessage(http.StatusUnauthorized, err.Error(), c)
		return
	}
	response.WithDetail(http.StatusCreated, res, "登录成功", c)
}

func AddManager(c *gin.Context) {

}

func DeleteManager(c *gin.Context) {

}

// @Tags manager
// @Summary 管理员密码修改
// @Produce application/json
// @Param data body request.ChangePasswd true "id, passwd, passwdNew"
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"修改成功"}"
// @Router /v1/manager/:id/passwd [put]
func ChangePasswd(c *gin.Context) {
	var req request.ChangePasswd
	_ = c.ShouldBindJSON(&req)
	id, _ := strconv.Atoi(c.Param("id"))
	req.ID = uint(id)
	m := &model.Manager{
		AM_MODEL: model.AM_MODEL{
			ID: req.ID,
		},
		Password: req.Password,
	}
	_, err := service.ChangePasswd(m, req.NewPasswd)
	if err != nil {
		response.WithMessage(http.StatusUnauthorized, "修改失败，密码验证未通过", c)
		return
	}
	response.WithMessage(http.StatusOK, "修改成功", c)
}

// @Tags manager
// @Summary 获取管理员列表
// @Produce application/json
// @Success 200 {string} json "{"success":true, "data":{}, "msg":""}"
// @Router /v1/manager [get]
func GetManagerList(c *gin.Context) {
	managers, _ := service.GetManagerList()
	response.WithDetail(http.StatusOK, managers, "查询成功", c)
}
