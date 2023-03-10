package util

import (
	"errors"
	"github.com/DodiNCer/tiktok/biz/model/user_gorm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var JwtKey = []byte("tiktok")

type MyClaims struct {
	UserName string `json:"name,omitempty"`
	UserId   int64  `json:"id,omitempty"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(userName string, userId int64, expireTime time.Time) (string, error) {
	SetClaims := MyClaims{
		UserName: userName,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tiktok",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	if token == "" {
		return &MyClaims{}, errors.New("token不存在")
	}
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	key, _ := setToken.Claims.(*MyClaims)
	if setToken.Valid {
		return key, nil
	}
	if !setToken.Valid {
		return nil, errors.New("token错误")

	}
	if time.Now().Unix() > key.ExpiresAt {
		return key, errors.New("token过期")
	}
	return key, nil
}

// JWT 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		_, err := CheckToken(token)
		if err != nil {
			c.JSON(http.StatusOK, user_gorm.BaseResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
