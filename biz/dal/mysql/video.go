package mysql

import (
	"github.com/DodiNCer/tiktok/biz/model"
	"time"
)

// 查询用户投稿视频数
func QueryVideoNumFromUser(uid int64) int64 {
	var count int64
	db := DB.Model(&model.Video{})
	db.Where("video.creator_id = ? AND is_deleted != TRUE", uid).Count(&count)
	return count
}

// 查找视频信息
func QueryVideos(videoId int64) (*model.Video, error) {
	var video *model.Video
	err := DB.Where("id = ? AND is_deleted = ?", videoId, 0).Find(&video).Error
	return video, err
}

func QueryVideoList(creatorId int64) ([]*model.Video, error) {
	var videoList []*model.Video
	err := DB.Where("creator_id = ? AND is_deleted = ?", creatorId, 0).Find(&videoList).Error
	return videoList, err
}

func CreatVideo(userId int64, title string, path string, coverPath string) error {
	err := DB.Create(&model.Video{
		Title:      title,
		Path:       path,
		CreatorId:  userId,
		CreateTime: time.Now(),
		CoverPath:  coverPath,
		IsDeleted:  false,
		UpdateTime: time.Now(),
	}).Error
	return err
}

func QueryVideoByTime(lastTime int64) ([]*model.Video, error) {
	var videoRes []*model.Video
	var err error
	if lastTime != 0 {
		err = DB.Model(model.Video{}).Where("create_time <= ?", time.Unix(lastTime, 0)).Order("create_time desc").Limit(30).Find(&videoRes).Error
	} else {
		db := DB.Model(model.Video{})
		db = db.Order("create_time desc")
		db = db.Limit(30).Find(&videoRes)
	}
	return videoRes, err
}
