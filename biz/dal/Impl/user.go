package Impl

import (
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/DodiNCer/tiktok/util"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"time"
)

// Register 用户注册
func Register(c *app.RequestContext) (model.UserResponse, error) {
	var userResponse model.UserResponse
	var token string
	userName := c.Query("username")
	password := c.Query("password")
	userId, err := mysql.Register(userName, password)
	if err != nil {
		return userResponse, err
	}
	token, err = util.SetToken(userName, userId, time.Now().Add(time.Hour*240))
	if err != nil {
		return userResponse, err
	}
	userResponse.UserId = userId
	userResponse.Token = token

	return userResponse, nil
}

// Login 用户登录
func Login(c *app.RequestContext) (model.UserResponse, error) {
	var userResponse model.UserResponse
	var token string
	username := c.Query("username")
	password := c.Query("password")
	userId, err := mysql.Login(username, password)
	if err != nil {
		return userResponse, err
	}
	token, err = util.SetToken(username, userId, time.Now().Add(time.Hour*240))
	if err != nil {
		return userResponse, err
	}
	userResponse.UserId = userId
	userResponse.Token = token
	return userResponse, nil
}

// UserInfo 查询用户信息
func UserInfo(c *app.RequestContext) (model.UserInfoResponse, error) {
	var userInfoResponse model.UserInfoResponse
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
