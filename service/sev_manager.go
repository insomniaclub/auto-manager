package service

import (
	"auto-manager/global"
	"auto-manager/model"
	"auto-manager/model/response"
	"auto-manager/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID         uint
	Nickname   string
	Password   string
	BufferTime int64
	jwt.StandardClaims
}

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
	loginRes := response.Login{
		Manager:   manager,
		Token:     tokenString,
		ExpiresAt: expiresTime,
	}
	return &loginRes, nil
}

func token(m *model.Manager) (int64, string, error) {
	claims := Claims{
		ID:         m.ID,
		Nickname:   m.Nickname,
		BufferTime: global.AM_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Unix() + global.AM_CONFIG.JWT.ExpiresTime,
			Issuer:    global.AM_CONFIG.JWT.SigningKey,
			NotBefore: time.Now().Unix() - 1000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("api.autoManager"))
	return claims.StandardClaims.ExpiresAt * 1000, tokenString, err
}
