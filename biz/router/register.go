// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	favorite_gorm "github.com/DodiNCer/tiktok/biz/router/favorite_gorm"
	comment_gorm "github.com/DodiNCer/tiktok/biz/router/comment_gorm"
	follower_gorm "github.com/DodiNCer/tiktok/biz/router/follower_gorm"
	publish_gorm "github.com/DodiNCer/tiktok/biz/router/publish_gorm"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!

	publish_gorm.Register(r)

	favorite_gorm.Register(r)

	comment_gorm.Register(r)

	follower_gorm.Register(r)
}
