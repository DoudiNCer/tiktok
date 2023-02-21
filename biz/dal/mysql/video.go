package mysql

import "github.com/DodiNCer/tiktok/biz/model"

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
