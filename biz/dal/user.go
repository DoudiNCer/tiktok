package dal

import (
	"github.com/DodiNCer/tiktok/Impl"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// Register 用户注册
func Register(c *app.RequestContext) {
	var response model.UserResponse
	var err error
	response, err = Impl.Register(c)
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
func Login(c *app.RequestContext) {
	var response model.UserResponse
	var err error
	response, err = Impl.Login(c)
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
func UserInfo(c *app.RequestContext) {
	var response model.UserInfoResponse
	var err error
	response, err = Impl.UserInfo(c)
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
