// Code generated by hertz generator.

package register

import (
	handler "github.com/DodiNCer/tiktok/biz/handler"
	"github.com/DodiNCer/tiktok/biz/handler/user_gorm"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func CustomizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	// your code ...

	//用户注册
	r.POST("/douyin/user/register", user_gorm.Register)
	//用户登录
	r.POST("/douyin/user/login/", user_gorm.Login)
	//用户信息
	r.GET("/douyin/user/", user_gorm.UserInfo)
}
