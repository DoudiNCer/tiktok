package user_gorm

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/model/user_gorm"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// Register 用户注册
func Register(ctx context.Context, c *app.RequestContext) {
	var response user_gorm.UserResponse
	var err error
	response, err = RegisterImpl(c)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "注册失败:" + err.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "注册成功"
	c.JSON(http.StatusOK, response)
}

// Login 用户登录
func Login(ctx context.Context, c *app.RequestContext) {
	var response user_gorm.UserResponse
	var err error
	response, err = LoginImpl(c)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "登录失败:" + err.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "登录成功"
	c.JSON(http.StatusOK, response)
	return
}

// UserInfo 用户信息
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var response user_gorm.UserInfoResponse
	var err error
	response, err = UserInfoImpl(c)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "查询失败:" + err.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "查询成功"
	c.JSON(http.StatusOK, response)
	return
}
