package user_gorm

import (
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	"github.com/DodiNCer/tiktok/biz/model/user_gorm"
	"github.com/DodiNCer/tiktok/biz/util"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"time"
)

// RegisterImpl Register 用户注册
func RegisterImpl(c *app.RequestContext) (user_gorm.UserResponse, error) {
	var userResponse user_gorm.UserResponse
	var token string
	userName := c.Query("username")
	password := c.Query("password")
	userId, err := mysql.Register(userName, password)
	if err != nil {
		return userResponse, err
	}
	token, err = util.SetToken(userName, userId, time.Now().Add(time.Hour*360))
	if err != nil {
		return userResponse, err
	}
	userResponse.UserId = userId
	userResponse.Token = token

	return userResponse, nil
}

// LoginImpl Login 用户登录
func LoginImpl(c *app.RequestContext) (user_gorm.UserResponse, error) {
	var userResponse user_gorm.UserResponse
	var token string
	username := c.Query("username")
	password := c.Query("password")
	userId, err := mysql.Login(username, password)
	if err != nil {
		return userResponse, err
	}
	token, err = util.SetToken(username, userId, time.Now().Add(time.Hour*360))
	if err != nil {
		return userResponse, err
	}
	userResponse.UserId = userId
	userResponse.Token = token
	return userResponse, nil
}

// UserInfoImpl UserInfo 查询用户信息
func UserInfoImpl(c *app.RequestContext) (user_gorm.UserInfoResponse, error) {
	var userInfoResponse user_gorm.UserInfoResponse
	token := c.Query("token")
	key, err := util.CheckToken(token)
	if err != nil {
		return userInfoResponse, err
	}
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	userInfoResponse, err = mysql.UserInfo(userId, key.UserId)
	if err != nil {
		return userInfoResponse, err
	}
	return userInfoResponse, nil
}
