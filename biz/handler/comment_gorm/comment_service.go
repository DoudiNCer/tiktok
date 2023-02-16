// Code generated by hertz generator.

package comment_gorm

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/DodiNCer/tiktok/biz/model/comment_gorm"
	"github.com/DodiNCer/tiktok/biz/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

// CreateComment .
// @router /douyin/comment/action/ [POST]
func CreateComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment_gorm.CommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
	}

	resp := new(comment_gorm.CommentActionResponse)

	token := req.Token
	key, err := util.CheckToken(token)
	if err != nil {
		c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_RTErr, StatusMsg: err.Error()})
		return
	}
	userId := key.UserId

	if req.ActionType == 1 {
		comment, err := mysql.CreateComment(&model.Comment{CreatorUid: userId, Text: req.CommentText, VideoId: req.VideoID, CreatedAt: time.Now()})
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//查找用户名称
		user, err := mysql.QueryUserByUid(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//查找关注总数
		_, followTotal, err := mysql.QueryFollow(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//查找粉丝总数
		_, followerTotal, err := mysql.QueryFollower(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//查看是否关注自己
		self, err := mysql.QueryFollowSelf(userId)
		if err != nil {
			return
		}

		//封装用户响应数据
		userResp := comment_gorm.User{ID: userId, Name: user.Name, FollowCount: followTotal, FollowerCount: followerTotal, IsFollow: self == 1}

		//封装评论响应数据
		commentResp := comment_gorm.Comment{ID: comment.Id, User: &userResp, Content: comment.Text, CreateDate: comment.CreatedAt.String()}

		resp.StatusCode = comment_gorm.Code_Success
		resp.StatusMsg = "success"
		resp.Comment = &commentResp

	} else if req.ActionType == 2 {
		err := mysql.DeleteComment(req.CommentID)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		resp.StatusCode = comment_gorm.Code_Success
		resp.StatusMsg = "success"
	} else {
		c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
	}

	c.JSON(consts.StatusOK, resp)
}

// QueryCommentList .
// @router /douyin/comment/list/ [GET]
func QueryCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment_gorm.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(comment_gorm.CommentListResponse)

	c.JSON(consts.StatusOK, resp)
}