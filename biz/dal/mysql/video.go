package mysql

import "github.com/DodiNCer/tiktok/biz/model"

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
