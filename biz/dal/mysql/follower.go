package mysql

import (
	"errors"
	"github.com/DodiNCer/tiktok/biz/model"
	"time"
)

func CreateFollower(followers []*model.Follower) error {
	return DB.Create(followers).Error
}

func UpdateFollower(uid, toUid int64, actionType int32) error {
	db := DB.Model(model.Follower{})
	db = db.Where("user_uid = ?", uid).Where("to_user_uid = ?", toUid)
	var status int32
	if actionType == 1 {
		status = 0
	} else if actionType == 2 {
		status = 1
	} else {
		return errors.New("非法参数actionType")
	}
	if err := db.Limit(1).Updates(map[string]interface{}{
		"is_deleted":  status,
		"update_time": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

// QueryFollow 查询关注
func QueryFollow(uid int64) ([]*model.Follower, int64, error) {
	db := DB.Model(model.Follower{})
	db = db.Where("user_uid = ?", uid).Where("is_deleted = 0")

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
	db = db.Where("to_user_uid = ?", uid).Where("is_deleted = 0")

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

func QueryForCheck(uid, toUid int64) (*model.Follower, int64, error) {
	db := DB.Model(model.Follower{})
	db = db.Where("user_uid = ?", uid).Where("to_user_uid = ?", toUid)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var res *model.Follower
	if err := db.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

// QueryIfFollowSomeone 查询是否存在关注的关系
func QueryIfFollowSomeone(commentCreatorId int64, userID int64) (int64, error) {
	db := DB.Model(model.Follower{})
	db.Where(model.Follower{ToUserUid: commentCreatorId, UserUid: userID}).Where("is_deleted = 0")
	var total int64
	err := db.Count(&total).Error
	return total, err
}
