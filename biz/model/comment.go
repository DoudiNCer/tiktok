package model

import (
	"time"
)

type Comment struct {
	Id         int64 `gorm:"primaryKey"`
	CreatorUid int64
	Text       string
	VideoId    int64
	CreatedAt  time.Time `gorm:"column:create_time"`
	UpdatedAt  time.Time `gorm:"column:update_time"`
	IsDeleted  int8
}

func (f *Comment) TableName() string {
	return "comment"
}
