package v1

import (
	"auto-manager/model"
	"auto-manager/model/response"
	"auto-manager/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags position
// @Summary 新增RFID与地址关系
// @Produce application/json
// @Param data body model.Position true "-, rfid_id, position"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/position [post]
func AddPosition(c *gin.Context) {
	var p model.Position
	_ = c.ShouldBindJSON(&p)
	position, err := service.AddPosition(&p)
	if err != nil {
		response.WithMessage(http.StatusInternalServerError, "新增记录失败", c)
		return
	}
	response.WithDetail(http.StatusCreated, position, "新增记录成功", c)
}

// @Tags position
// @Summary 删除RFID与地址关系
// @Produce application/json
// @Success 204 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/position/:id [delete]
func DeletePosition(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	position, err := service.DeletePosition(uint(id))
	if err != nil {
		response.WithMessage(http.StatusInternalServerError, "删除记录失败", c)
		return
	}
	response.WithDetail(http.StatusNoContent, position, "删除记录成功", c)
}

// @Tags position
// @Summary 修改RFID与地址关系
// @Produce application/json
// @Param data body model.Position true "-, rfid_id, position"
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/position/:id [put]
func ModifyPosition(c *gin.Context) {
	var p model.Position
	_ = c.ShouldBindJSON(&p)
	id, _ := strconv.Atoi(c.Param("id"))
	p.ID = uint(id)
	_, err := service.ModifyPosition(&p)
	if err != nil {
		response.WithMessage(http.StatusInternalServerError, "修改记录失败", c)
		return
	}
	response.WithDetail(http.StatusOK, p, "修改记录成功", c)
}

// @Tags position
// @Summary 获取RFID与地址对应表
// @Produce application/json
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/position [get]
func GetPositionList(c *gin.Context) {
	positions, err := service.GetPositionList()
	if err != nil {
		response.WithMessage(http.StatusNotFound, "查找失败", c)
		return
	}
	response.WithDetail(http.StatusOK, positions, "查找成功", c)
}
