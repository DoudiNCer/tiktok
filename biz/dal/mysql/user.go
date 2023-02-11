package mysql

import (
	"github.com/DodiNCer/tiktok/biz/model"
)

func QueryUserByUid(uid int64) (*model.User, error) {

	db := DB.Model(model.User{})

	db = db.Where("id = ?", uid)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}
	var res *model.User
	if err := db.First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}