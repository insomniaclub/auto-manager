package service

import (
	"auto-manager/global"
	"auto-manager/model"
	"auto-manager/utils"
)

func Login(m *model.Manager) (error, *model.Manager) {
	var manager model.Manager
	m.Password = utils.MD5Encrypt(m.Password)
	err := global.AM_DB.Table("manager").Where("nickname = ? AND password = ?", m.Nickname, m.Password).First(&manager).Error
	return err, &manager
}
