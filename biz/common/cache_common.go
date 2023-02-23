package common

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var CacheManager = cache.New(24*time.Hour, time.Duration(time.Monday))
var KeyAddUser = "key_add_user"
var KeyAddFriend = "key_add_friend"

// DeleteUserReferTo 改变了user信息的接口都需要执行此操作
// 例如：赞，关注，更新用户信息
func DeleteUserReferTo(uid string) {
	CacheManager.Delete(uid + KeyAddUser)
	CacheManager.Delete(uid + KeyAddFriend)
}
