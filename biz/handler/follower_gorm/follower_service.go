// Code generated by hertz generator.

package follower_gorm

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	"github.com/DodiNCer/tiktok/biz/model"
	follower_gorm "github.com/DodiNCer/tiktok/biz/model/follower_gorm"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"time"
)

// CreateFollower .
// @router /douyin/relation/action/ [POST]
func CreateFollower(ctx context.Context, c *app.RequestContext) {
	var err error
	var req follower_gorm.CreateFollowerRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	//从请求中获取to_user_id
	toUserUid := req.ToUserID
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	//在follower表中创建记录
	if err = mysql.CreateFollower([]*model.Follower{
		{
			ToUserUid:  toUserUid,
			UserUid:    2,
			CreateTime: time.Now(),
			IsDeleted:  false,
			UpdateTime: time.Now(),
		},
	}); err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(follower_gorm.CreateFollowerResponse)
	resp.StatusMsg = "关注成功"

	c.JSON(consts.StatusOK, resp)
}

// QueryFollowList .
// @router /douyin/relatioin/follow/list/ [GET]
func QueryFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req follower_gorm.QueryFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//从请求获取uid
	userId := req.UserID
	parseInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return
	}
	//查询关注列表
	followList, _, err := mysql.QueryFollow(parseInt)
	if err != nil {
		return
	}

	resp := new(follower_gorm.QueryFollowListResponse)
	resp.StatusMsg = "请求成功"

	var userList = resp.GetUserList()

	for i := 0; i < len(followList); i++ {
		//创建载体对象
		var userSingle follower_gorm.User
		//查询出的关注对象
		follower := followList[i]
		//获取关注对象uid
		uid := follower.ToUserUid

		//查询关注对象的关注总数
		_, followCount, err := mysql.QueryFollow(uid)
		if err != nil {
			return
		}
		//查询关注对象的粉丝总数
		_, followerCount, err := mysql.QueryFollower(uid)
		if err != nil {
			return
		}
		//查询关注对象信息
		user, err := mysql.QueryUserByUid(uid)
		if err != nil {
			return
		}
		//数据装配
		userSingle.IsFollow = true
		userSingle.Name = user.Name
		userSingle.ID = uid
		userSingle.FollowCount = followCount
		userSingle.FollowerCount = followerCount
		userList = append(userList, &userSingle)
	}
	resp.UserList = userList

	c.JSON(consts.StatusOK, resp)
}
