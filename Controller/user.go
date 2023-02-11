package Controller

import (
	"github.com/DodiNCer/tiktok/Dao"
	"github.com/DodiNCer/tiktok/Impl"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// Register 用户注册
func Register(c *app.RequestContext) {
	var response Dao.UserResponse
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
