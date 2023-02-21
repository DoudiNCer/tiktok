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

		if len(req.CommentText) == 0 {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_ParamInvalid, StatusMsg: "不能发空字符串"})
			return
		}

		comment, err := mysql.CreateComment(&model.Comment{CreatorUid: userId, Text: req.CommentText, VideoId: req.VideoID, CreatedAt: time.Now(), UpdatedAt: time.Now()})
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
		self, err := mysql.QueryIfFollowSomeone(userId, userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取用户的总点赞数
		favoriteCount, err := mysql.QueryFavoriteCount(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取该用户作品的被点赞总数数
		totalFavorited, err := mysql.QueryTotalFavorited(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取用户的作品及作品数量
		workCount, err := mysql.QueryWorkCount(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取用户头像
		portraitPath, err := mysql.QueryPortraitPathByUserId(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取用户个人页顶部大图
		backgroundImage, err := mysql.QueryBackgroundImageByUserId(userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: "获取用户个人页顶部大图失败"})
			return
		}

		//获取用户个人简介
		signature, err := mysql.QuerySignatureByUserId(userId)
		if err != nil {
			return
		}

		//封装用户响应数据
		userResp := comment_gorm.User{ID: userId, Name: user.Name, FollowCount: followTotal, FollowerCount: followerTotal, IsFollow: self == 1, FavoriteCount: favoriteCount, WorkCount: workCount, TotalFavorited: totalFavorited,
			Avatar: portraitPath, Signature: signature, BackgroundImage: backgroundImage,
		}

		//封装评论响应数据
		commentResp := comment_gorm.Comment{ID: comment.Id, User: &userResp, Content: comment.Text, CreateDate: comment.CreatedAt.Format("01-02")}

		resp.Comment = &commentResp

	} else if req.ActionType == 2 {
		err := mysql.DeleteComment(req.CommentID)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

	} else {
		c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
	}

	resp.StatusCode = comment_gorm.Code_Success
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// QueryCommentList .
// @router /douyin/comment/list/ [GET]
func QueryCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment_gorm.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
		return
	}

	//获取用户Id
	token := req.Token
	key, err := util.CheckToken(token)
	if err != nil {
		c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_RTErr, StatusMsg: err.Error()})
		return
	}
	userId := key.UserId

	//根据videoId查找对应源评论集
	videoId := req.VideoID
	comments, err := mysql.QueryCommentsByVideoId(videoId)
	if err != nil {
		c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
		return
	}

	//构造返回的评论集
	commentsResp := make([]*comment_gorm.Comment, 0)

	for i := 0; i < len(comments); i++ {

		comment := comments[i]

		//查找评论用户名称
		user, err := mysql.QueryUserByUid(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//查找评论者关注总数
		_, followTotal, err := mysql.QueryFollow(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//查找评论者粉丝总数
		_, followerTotal, err := mysql.QueryFollower(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//查看评论者和用户是否存在关注的关系
		isFollow, err := mysql.QueryIfFollowSomeone(comment.CreatorUid, userId)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取评论者的总点赞数
		favoriteCount, err := mysql.QueryFavoriteCount(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取该评论者作品的被点赞总数数
		totalFavorited, err := mysql.QueryTotalFavorited(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取评论者的作品及作品数量
		workCount, err := mysql.QueryWorkCount(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取评论者头像
		portraitPath, err := mysql.QueryPortraitPathByUserId(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: err.Error()})
			return
		}

		//获取评论者个人页顶部大图
		backgroundImage, err := mysql.QueryBackgroundImageByUserId(comment.CreatorUid)
		if err != nil {
			c.JSON(200, &comment_gorm.CommentActionResponse{StatusCode: comment_gorm.Code_DBErr, StatusMsg: "获取用户个人页顶部大图失败"})
			return
		}

		//获取评论者个人简介
		signature, err := mysql.QuerySignatureByUserId(comment.CreatorUid)
		if err != nil {
			return
		}

		//封装评论者响应数据
		userResp := comment_gorm.User{ID: comment.CreatorUid, Name: user.Name, FollowCount: followTotal, FollowerCount: followerTotal, IsFollow: isFollow == 1, FavoriteCount: favoriteCount, WorkCount: workCount, TotalFavorited: totalFavorited,
			Avatar: portraitPath, Signature: signature, BackgroundImage: backgroundImage,
		}

		//封装返回的评论
		var commentResp comment_gorm.Comment
		commentResp.ID = comment.Id
		commentResp.User = &userResp
		commentResp.Content = comment.Text
		commentResp.CreateDate = comment.CreatedAt.Format("01-02")

		//封装返回的评论集
		commentsResp = append(commentsResp, &commentResp)
	}

	resp := new(comment_gorm.CommentListResponse)
	resp.CommentList = commentsResp
	resp.StatusCode = comment_gorm.Code_Success
	resp.StatusMsg = "success"

	c.JSON(consts.StatusOK, resp)
}
