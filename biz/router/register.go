// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	follower_gorm "github.com/DodiNCer/tiktok/biz/router/follower_gorm"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	follower_gorm.Register(r)

}