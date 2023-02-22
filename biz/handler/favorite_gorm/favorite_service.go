// Code generated by hertz generator.

package favorite_gorm

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/DodiNCer/tiktok/biz/model/follower_gorm"
	"github.com/DodiNCer/tiktok/biz/util"
	"gorm.io/gorm"
	"time"

	"github.com/DodiNCer/tiktok/biz/model/favorite_gorm"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite_gorm.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{StatusCode: follower_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
		return
	}

	//获取相关参数
	videoId := req.VideoID
	actionType := req.ActionType
	token := req.Token
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_RTErr,
			StatusMsg:  err.Error(),
		})
	}

	//从token中拿取uid
	key, err := util.CheckToken(token)
	userId := key.UserId
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_RTErr,
			StatusMsg:  err.Error(),
		})
	}

	//数据库操作
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		var favorite model.Favorite

		//查找是否存在该条点赞记录
		if err = tx.Where(map[string]interface{}{
			"creator_id": userId,
			"video_id":   videoId,
			"is_deleted": 0,
		}).Limit(1).Find(&favorite).Error; err != nil {
			return err
		}

		//点赞和取消点赞操作
		if favorite.Id != 0 && actionType == 2 {
			//取消点赞操作：去除点赞记录
			if err = tx.Model(&favorite).Limit(1).Updates(map[string]interface{}{
				"is_deleted":  1,
				"update_time": time.Now(),
			}).Error; err != nil {
				return err
			}
		} else if favorite.Id == 0 && actionType == 1 {
			//点赞操作：创建点赞记录
			if err = tx.Create(&model.Favorite{
				CreatorId:  userId,
				VideoId:    videoId,
				IsDeleted:  false,
				UpdateTime: time.Now(),
				CreateTime: time.Now(),
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
	}

	resp := new(favorite_gorm.FavoriteActionResponse)
	resp.StatusCode = follower_gorm.Code_Success
	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [POST]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite_gorm.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{StatusCode: follower_gorm.Code_ParamInvalid, StatusMsg: err.Error()})
		return
	}

	reqUserId := req.UserID
	token := req.Token

	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_RTErr,
			StatusMsg:  err.Error(),
		})
	}

	//从token中拿取uid
	key, err := util.CheckToken(token)
	userId := key.UserId
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_RTErr,
			StatusMsg:  err.Error(),
		})
	}
	//验证参数里的id和token的id一致
	if reqUserId != userId {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_RTErr,
			StatusMsg:  err.Error(),
		})
	}

	var videoList []*favorite_gorm.Video

	//数据库操作
	err = func() error {
		var favorites []*model.Favorite
		if favorites, err = mysql.QueryFavoritesByCreatorId(userId); err != nil {
			return err
		}

		var video *model.Video
		var user *model.User
		var followCount int64   //关注总数
		var followerCount int64 //粉丝总数
		var favoriteCount int64 //视频点赞总数
		var commentCount int64  //视频评论总数
		var isbool int64        //用户是否关注视频作者

		//循环查找视频数据
		for _, favorite := range favorites {
			//查找视频信息
			if video, err = mysql.QueryVideos(favorite.VideoId); err != nil {
				return err
			}
			//查找视频作者信息
			if user, err = mysql.QueryUserByUid(video.CreatorId); err != nil {
				return err
			}
			//查找视频作者关注总数
			if _, followCount, err = mysql.QueryFollow(user.Id); err != nil {
				return err
			}
			//查找视频作者粉丝总数
			if _, followerCount, err = mysql.QueryFollower(user.Id); err != nil {
				return err
			}
			//查找视频点赞总数
			if favoriteCount, err = mysql.QueryFavoriteNumByVideo(favorite.VideoId); err != nil {
				return err
			}
			//查找视频评论总数
			if commentCount, err = mysql.QueryCommentCountByVideo(favorite.VideoId); err != nil {
				return err
			}
			//查找用户是否关注视频作者
			if isbool, err = mysql.QueryIfFollowSomeone(user.Id, userId); err != nil {
				return err
			}
			//拼装数据
			videoList = append(videoList, &favorite_gorm.Video{
				ID: video.Id,
				Author: &follower_gorm.User{
					ID:            user.Id,
					Name:          user.Name,
					FollowCount:   followCount,
					FollowerCount: followerCount,
					IsFollow: func() bool {
						if isbool == 1 {
							return true
						} else {
							return false
						}
					}(),
				},
				PlayURL:       video.Path,
				CoverURL:      video.CoverPath,
				FavoriteCount: favoriteCount,
				CommentCount:  commentCount,
				IsFavorite:    true,
				Title:         video.Title,
			})
		}
		return nil
	}()
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
	}

	resp := new(favorite_gorm.FavoriteListResponse)
	resp.VideoList = videoList
	c.JSON(consts.StatusOK, resp)
}
