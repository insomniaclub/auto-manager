package request

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID         uint
	Nickname   string
	Password   string
	BufferTime int64
	jwt.StandardClaims
}

type Login struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Register struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type ChangePasswd struct {
	ID        uint   `json:"id"`
	Password  string `json:"password"`
	NewPasswd string `json:"newPasswd"`
}
