// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	follower_gorm "github.com/DodiNCer/tiktok/biz/router/follower_gorm"
	friend_list_gorm "github.com/DodiNCer/tiktok/biz/router/friend_list_gorm"
	ws_chat "github.com/DodiNCer/tiktok/biz/router/ws_chat"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	friend_list_gorm.Register(r)

	ws_chat.Register(r)

	follower_gorm.Register(r)
}
