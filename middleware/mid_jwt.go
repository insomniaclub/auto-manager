package middleware

import (
	"auto-manager/global"
	"auto-manager/model/request"
	"auto-manager/model/response"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenMalformed   = errors.New("not a token")
	TokenNotValidYet = errors.New("token not active yet")
	TokenInvalid     = errors.New("token invalid")
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.WithDetail(
				http.StatusUnauthorized, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := &JWT{
			[]byte(global.AM_CONFIG.JWT.SigningKey),
		}
		claims, err := j.ParseToken(token)
		if err != nil {
			switch err {
			case TokenExpired:
				response.WithDetail(
					http.StatusUnauthorized, gin.H{"reload": true}, "授权以过期", c)
				c.Abort()
				return
			default:
				response.WithDetail(
					http.StatusUnauthorized, gin.H{"reload": true}, err.Error(), c)
				c.Abort()
				return
			}
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.AM_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func (j *JWT) CreateToken(claims request.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, TokenMalformed
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, TokenExpired
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, TokenNotValidYet
			default:
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
