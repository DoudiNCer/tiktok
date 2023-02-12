package mysql

import (
	"github.com/DodiNCer/tiktok/biz/model"
)

func CreateFollower(followers []*model.Follower) error {
	return DB.Create(followers).Error
}

func UpdateFollower(followers *model.Follower) error {
	return DB.Updates(followers).Error
}

// QueryFollow 查询关注
func QueryFollow(uid int64) ([]*model.Follower, int64, error) {
	db := DB.Model(model.Follower{})
	db = db.Where("user_uid = ?", uid)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var res []*model.Follower
	if err := db.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

// QueryFollower 查询粉丝
func QueryFollower(uid int64) ([]*model.Follower, int64, error) {
	db := DB.Model(model.Follower{})
	db = db.Where("to_user_uid = ?", uid)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var res []*model.Follower
	if err := db.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, total, nil
}
