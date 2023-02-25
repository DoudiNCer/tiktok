// Code generated by hertz generator.

package feed_gorm

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/common"
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/DodiNCer/tiktok/biz/model/favorite_gorm"
	"github.com/DodiNCer/tiktok/biz/model/follower_gorm"
	"github.com/DodiNCer/tiktok/biz/util"
	"github.com/patrickmn/go-cache"
	"strconv"

	feed_gorm "github.com/DodiNCer/tiktok/biz/model/feed_gorm"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// QueryFeedList .
// @router /douyin/feed_gorm/ [GET]
func QueryFeedList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed_gorm.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	time := req.LastTime
	token := req.Token

	var uid = int64(0)
	if token != "" {
		checkToken, err := util.CheckToken(token)
		if err != nil {
			c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
				StatusCode: follower_gorm.Code_DBErr,
				StatusMsg:  err.Error(),
			})
			return
		}
		uid = checkToken.UserId
	}

	var videoListRes []*favorite_gorm.Video

	var videoList []*model.Video
	var user *model.User
	var followCount int64
	var followerCount int64
	var favoriteCount int64
	var commentCount int64
	var isbool int64
	//数据库操作
	err = func() error {

		videoList, err = mysql.QueryVideoByTime(time)
		if err != nil {
			return err
		}
		for _, video := range videoList {
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
			userRes := &follower_gorm.User{
				ID:            user.Id,
				Name:          user.Name,
				FollowCount:   followCount,
				FollowerCount: followerCount,
				IsFollow:      false,
			}

			if favoriteCount, err = mysql.QueryFavoriteNumByVideo(video.Id); err != nil {
				return err
			}
			if commentCount, err = mysql.QueryCommentCountByVideo(video.Id); err != nil {
				return err
			}

			if token != "" {
				if v, found := common.CacheManager.Get(strconv.FormatInt(uid, 10) + strconv.FormatInt(video.Id, 10) + common.KeyFavoriteIs); found == true {
					isbool = v.(int64)
				} else if isbool, err = mysql.QueryFavoriteIs(uid, video.Id); err != nil {
					return err
				} else {
					common.CacheManager.Set(strconv.FormatInt(uid, 10)+strconv.FormatInt(video.Id, 10)+common.KeyFavoriteIs, isbool, cache.DefaultExpiration)
				}
			} else {
				isbool = 0
			}
			videoListRes = append(videoListRes, &favorite_gorm.Video{
				ID:            video.Id,
				Author:        userRes,
				PlayURL:       video.Path,
				CoverURL:      video.CoverPath,
				FavoriteCount: favoriteCount,
				CommentCount:  commentCount,
				IsFavorite: func() bool {
					if isbool == 1 {
						return true
					} else {
						return false
					}
				}(),
				Title: video.Title,
			})
		}

		return nil
	}()
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
		return
	}

	resp := new(feed_gorm.FeedResponse)
	resp.VideoList = videoListRes

	resp.StatusCode = 0

	resp.StatusMsg = "ok"

	c.JSON(consts.StatusOK, resp)
}
