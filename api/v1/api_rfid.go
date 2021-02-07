package v1

import "github.com/gin-gonic/gin"

// @Tags 待完成
// @Summary 新增记录
// @Produce application/json
// @Param data body model.RFID true "rfid与车辆id的映射"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/rfid [post]
func AddRecord(c *gin.Context) {

}

// @Tags 待完成
// @Summary 新增记录
// @Produce application/json
// @Param data body uint true "id"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/rfid/:id [delete]
func DeleteRecord(c *gin.Context) {

}

// @Tags 待完成
// @Summary 新增记录
// @Produce application/json
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/rfid [get]
func GetRecordList(c *gin.Context) {

}
