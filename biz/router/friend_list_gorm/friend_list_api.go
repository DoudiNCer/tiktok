// Code generated by hertz generator. DO NOT EDIT.

package FriendListGorm

import (
	friend_list_gorm "github.com/DodiNCer/tiktok/biz/handler/friend_list_gorm"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				{
					_list := _friend.Group("/list", _listMw()...)
					_list.GET("/", append(_getfriendlistMw(), friend_list_gorm.GetFriendList)...)
				}
			}
		}
	}
}
