package v1

// @Tags 待完成
// @Summary 新增记录
// @Produce application/json
// @Param data body model.RFID true "rfid与车辆id的映射"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/car [post]
func AddRecord() {

}

// @Tags 待完成
// @Summary 新增记录
// @Produce application/json
// @Param data body uint true "id"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/car/:id [delete]
func DeleteRecord() {

}

// @Tags 待完成
// @Summary 新增记录
// @Produce application/json
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/car [get]
func GetRecordList() {

}
