package v1

// @Tags 待完成
// @Summary 微信小程序认证
// @Produce application/json
// @Param "暂定OpenID"
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"登录成功"}"
// @Router /v1/master/token [post]
func Checkin() {

}

// @Tags 待完成
// @Summary 注销账户
// @Produce application/json
// @Param ID "id"
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/master/:id [delete]
func CancelAccount() {

}

// @Tags 待完成
// @Summary 用户信息修改
// @Produce application/json
// @Param data body model.Master "完整的用户信息"
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/master/:id [put]
func ModifyMaster() {

}

// @Tags 待完成
// @Summary 获取用户列表
// @Produce application/json
// @Success 200 {string} json "{"success":true, "data":{}, "msg":"成功"}"
// @Router /v1/master [get]
func GetMasterList() {

}
