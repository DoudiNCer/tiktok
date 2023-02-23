// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	comment_gorm "github.com/DodiNCer/tiktok/biz/router/comment_gorm"
	favorite_gorm "github.com/DodiNCer/tiktok/biz/router/favorite_gorm"
	Feed "github.com/DodiNCer/tiktok/biz/router/feed_gorm"
	follower_gorm "github.com/DodiNCer/tiktok/biz/router/follower_gorm"
	friend_list_gorm "github.com/DodiNCer/tiktok/biz/router/friend_list_gorm"
	friend_talk_message_gorm "github.com/DodiNCer/tiktok/biz/router/friend_talk_message_gorm"
	publish_gorm "github.com/DodiNCer/tiktok/biz/router/publish_gorm"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	Feed.Register(r)

	friend_list_gorm.Register(r)

	friend_talk_message_gorm.Register(r)

	publish_gorm.Register(r)

	favorite_gorm.Register(r)

	comment_gorm.Register(r)

	follower_gorm.Register(r)
}
