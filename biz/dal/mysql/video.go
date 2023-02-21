package mysql

import "github.com/DodiNCer/tiktok/biz/model"

// 查询用户投稿视频数
func QueryVideoNumFromUser(uid int64) int64 {
	var count int64
	db := DB.Model(&model.Video{})
	db.Where("video.creator_id = ? AND is_deleted != TRUE", uid).Count(&count)
	return count
}
