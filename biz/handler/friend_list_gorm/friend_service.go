// Code generated by hertz generator.

package friend_list_gorm

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	friend_list_gorm "github.com/DodiNCer/tiktok/biz/model/friend_list_gorm"
	"github.com/DodiNCer/tiktok/biz/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// GetFriendList .
// @router /douyin/relation/friend/list/ [GET]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req friend_list_gorm.GetFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(200, &friend_list_gorm.GetFriendListResponse{StatusCode: friend_list_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
		return
	}

	_, err = util.CheckToken(req.Token) //校验token
	if err != nil {
		c.JSON(200, &friend_list_gorm.GetFriendListResponse{StatusCode: friend_list_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
		return
	}

	//从请求中获取to_user_id
	UserId := req.UserID
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		c.JSON(200, &friend_list_gorm.GetFriendListResponse{StatusCode: friend_list_gorm.Code_RTErr, StatusMsg: err.Error()})
		return
	}

	friends, err := mysql.QueryFriend(uid)
	if err != nil {
		c.JSON(200, &friend_list_gorm.GetFriendListResponse{
			StatusCode: friend_list_gorm.Code_RTErr,
			StatusMsg:  err.Error(),
		})
		return
	}

	resp := new(friend_list_gorm.GetFriendListResponse)
	userList := resp.GetUserList()

	for _, user := range friends {
		var friendSingle friend_list_gorm.FriendUser
		// 取出好友信息
		frienduser := user
		id := frienduser.Id
		name := frienduser.Name

		// 查询粉丝数
		followerNum, err := mysql.QueryFollowerNum(id)
		if err != nil {
			c.JSON(200, &friend_list_gorm.GetFriendListResponse{StatusCode: friend_list_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}
		// 查询关注数
		followNum, err := mysql.QueryFollowNum(id)
		if err != nil {
			c.JSON(200, &friend_list_gorm.GetFriendListResponse{StatusCode: friend_list_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}
		// 查询最新消息
		message, msgType, err := mysql.QueryLastMessage(uid, id)
		if err != nil {
			c.JSON(200, &friend_list_gorm.GetFriendListResponse{StatusCode: friend_list_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}
		// 填充数据
		friendSingle.ID = id
		friendSingle.Name = name
		friendSingle.IsFollow = true
		friendSingle.FollowCount = followNum
		friendSingle.FollowerCount = followerNum
		friendSingle.Message = message.Text
		friendSingle.MsgType = msgType
		userList = append(userList, &friendSingle)
	}

	c.JSON(consts.StatusOK, resp)
}