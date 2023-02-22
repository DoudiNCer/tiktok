package model

import (
	"time"
)

// User User实体类
type User struct {
	Id                    int64     `gorm:"column:id"`                      //用户id
	Name                  string    `gorm:"column:name"`                    // 用户名称
	Password              string    `gorm:"column:password"`                //用户密码
	CreateTime            time.Time `gorm:"column:create_time"`             //用户创建时间
	PortraitPath          string    `gorm:"column:portrait_path"`           //用户头像
	BackgroundPicturePath string    `gorm:"column:background_picture_path"` //用户个人页顶部大图
	Signature             string    `gorm:"column:signature"`               //个人简介
}

func (ui User) TableName() string {
	return "user"
}
