package v1

// @Tags 待完成
// @Summary 新增车辆
// @Produce application/json
// @Param data body model.Car true "完整Car信息"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/car [post]
func AddCar() {

}

// @Tags 待完成
// @Summary 注销车辆
// @Produce application/json
// @Param data body uint true "id"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/car/:id [delete]
func DeleteCar() {

}

// @Tags 待完成
// @Summary 修改车辆信息
// @Produce application/json
// @Param data body model.Car true "完整Car信息"
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/car/:id [put]
func ModifyCar() {

}

// @Tags 待完成
// @Summary 获取车辆列表
// @Produce application/json
// @Success 201 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/car [get]
func GetCarList() {

}
