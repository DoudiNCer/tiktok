package Impl

import (
	"github.com/DodiNCer/tiktok/Dao"
	"github.com/DodiNCer/tiktok/dbfunc"
	"github.com/DodiNCer/tiktok/util"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

func Register(c *app.RequestContext) (Dao.UserResponse, error) {
	var userResponse Dao.UserResponse
	var token string
	userName := c.Query("username")
	password := c.Query("password")
	userId, err := dbfunc.Register(userName, password)
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
