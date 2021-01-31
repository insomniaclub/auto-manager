package service

import (
	"auto-manager/global"
	"auto-manager/middleware"
	"auto-manager/model"
	"auto-manager/model/request"
	"auto-manager/model/response"
	"auto-manager/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Login(m *model.Manager) (*response.Login, error) {
	var manager model.Manager
	m.Password = utils.MD5Encrypt(m.Password)
	err := global.AM_DB.Table("manager").Where("nickname = ? AND password = ?", m.Nickname, m.Password).First(&manager).Error
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	expiresTime, tokenString, err := token(&manager)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	manager.Password = ""
	loginRes := response.Login{
		Manager:   manager,
		Token:     tokenString,
		ExpiresAt: expiresTime,
	}
	return &loginRes, nil
}

func ChangePasswd(m *model.Manager, newPasswd string) (*model.Manager, error) {
	var manager model.Manager
	m.Password = utils.MD5Encrypt(m.Password)
	err := global.AM_DB.Table("manager").Where("id = ? AND password = ?", m.ID, m.Password).First(&manager).Update("password", utils.MD5Encrypt(newPasswd)).Error
	return m, err
}

func GetManagerList() (*[]response.Manager, error) {
	var managers []response.Manager
	err := global.AM_DB.Table("manager").Find(&managers).Error
	return &managers, err
}

func token(m *model.Manager) (int64, string, error) {
	j := &middleware.JWT{SigningKey: []byte(global.AM_CONFIG.JWT.SigningKey)}
	claims := request.Claims{
		ID:         m.ID,
		Nickname:   m.Nickname,
		BufferTime: global.AM_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.AM_CONFIG.JWT.ExpiresTime,
			Issuer:    global.AM_CONFIG.JWT.SigningKey,
			NotBefore: time.Now().Unix() - 1000,
		},
	}
	token, err := j.CreateToken(claims)
	return claims.StandardClaims.ExpiresAt * 1000, token, err
}
