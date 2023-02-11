package mysql

import (
	"github.com/DodiNCer/tiktok/biz/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func QueryUserByUid(uid int64) (*model.User, error) {

	db := DB.Model(model.User{})

	if uid != 0 {
		db = db.Where("id = ?", uid)
	}
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
