package mysql

import "github.com/DodiNCer/tiktok/biz/model"

func QueryWorkCount(userId int64) (int64, error) {
	var count int64
	var work model.Video
	DB = DB.Model(&model.Video{}).Where("creator_id = ? AND is_deleted = 0", userId)
	err := DB.Count(&count).Error
	err = DB.Find(&work).Error
	return count, err
}
