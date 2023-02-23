// Code generated by hertz generator.

package publish_gorm

import (
	"bytes"
	"context"
	"github.com/DodiNCer/tiktok/biz/common"
	"github.com/DodiNCer/tiktok/biz/dal/mysql"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/DodiNCer/tiktok/biz/model/favorite_gorm"
	"github.com/DodiNCer/tiktok/biz/model/follower_gorm"
	"github.com/DodiNCer/tiktok/biz/model/publish_gorm"
	"github.com/DodiNCer/tiktok/biz/mw"
	"github.com/DodiNCer/tiktok/biz/mw/kitex_ffmpeg"
	"github.com/DodiNCer/tiktok/biz/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/patrickmn/go-cache"
	"strconv"
)

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish_gorm.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
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
		if user, err = mysql.QueryUserByUid(userId); err != nil {
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

		if videoList, err = mysql.QueryVideoList(userId); err != nil {
			return err
		}
		for _, video := range videoList {
			if favoriteCount, err = mysql.QueryFavoriteNumByVideo(video.Id); err != nil {
				return err
			}
			if commentCount, err = mysql.QueryCommentCountByVideo(video.Id); err != nil {
				return err
			}
			if v, found := common.CacheManager.Get(strconv.FormatInt(userId, 10) + strconv.FormatInt(video.Id, 10) + common.KeyFavoriteIs); found == true {
				isbool = v.(int64)
			} else if isbool, err = mysql.QueryFavoriteIs(userId, video.Id); err != nil {
				return err
			} else {
				common.CacheManager.Set(strconv.FormatInt(userId, 10)+strconv.FormatInt(video.Id, 10)+common.KeyFavoriteIs, isbool, cache.DefaultExpiration)
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
	}
	resp := new(publish_gorm.PublishListResponse)
	resp.VideoList = videoListRes
	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish_gorm.PublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	token := req.Token
	data := req.Data
	title := req.Title
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_RTErr,
			StatusMsg:  err.Error(),
		})
		return
	}

	key, err := util.CheckToken(token)

	id := key.UserId

	user, err := mysql.QueryUserByUid(id)
	if err != nil || user.Id == 0 {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
		return
	}

	//上传视频
	videoReader := bytes.NewReader(data)

	videoName, err := util.MinioUploadVideo(videoReader, videoReader.Size())
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
		return
	}

	workspace, err := mw.RPCClient.InitWorkspace(ctx, &kitex_ffmpeg.InitWorkspaceRequest{Whoami: "汤香b"})
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
		return
	}
	file := kitex_ffmpeg.File{FileName: "video", Content: data}
	files, err := mw.RPCClient.UploadFiles(context.Background(), &kitex_ffmpeg.UploadFilesRequest{Token: workspace.Token, Files: []*kitex_ffmpeg.File{&file}})
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
		return
	}
	execRes, err := mw.RPCClient.ExecFfmpeg(context.Background(), &kitex_ffmpeg.ExecRequest{Token: workspace.Token, Args_: []string{"-i", files.Files[0].FileID, "-ss", "00:00:02", "-frames:v", "1", "out.jpg"}})
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error() + "\nsout: " + execRes.Sout + "\nserr: " + execRes.Serr,
		})
		return
	}
	downloadFiles, err := mw.RPCClient.DownloadFiles(context.Background(), &kitex_ffmpeg.DownloadFilesRequest{Token: workspace.Token, FileIDs: []string{"a.out"}})
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error() + "\nsout: " + execRes.Sout + "\nserr: " + execRes.Serr,
		})
		return
	}
	//上传视频封面
	coverReader := bytes.NewReader(downloadFiles.Files[0].Content)
	_, err = mw.RPCClient.DropWorkspace(context.Background(), &kitex_ffmpeg.DropWorkspaceRequest{Token: workspace.Token, Whoami: "汤香香"})
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error() + "\nsout: " + execRes.Sout + "\nserr: " + execRes.Serr,
		})
		return
	}
	photo, err := util.MinioUploadPhoto(coverReader, coverReader.Size())
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
		return
	}

	//将相关信息写入数据库
	err = mysql.CreatVideo(id, title, mw.MinioLinkPrefix+videoName, mw.MinioLinkPrefix+photo)
	if err != nil {
		c.JSON(consts.StatusOK, &favorite_gorm.FavoriteActionResponse{
			StatusCode: follower_gorm.Code_DBErr,
			StatusMsg:  err.Error(),
		})
		return
	}
	//发布视频更新用户相关缓存
	common.DeleteUserReferTo(strconv.FormatInt(user.Id, 10))
	resp := new(publish_gorm.PublishActionResponse)
	resp.StatusCode = 0
	resp.StatusMsg = "上传视频成功"
	c.JSON(consts.StatusOK, resp)
}
