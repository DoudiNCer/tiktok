// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/DodiNCer/tiktok/biz/router/register"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Register registers all routers.
func Register(r *server.Hertz) {

	register.GeneratedRegister(r)

	register.CustomizedRegister(r)
}
