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
	IsDeleted  int8      //这里我没用bool值是因为会导致查询的时候会把软删的数据也给加进来
}

func (f *Comment) TableName() string {
	return "comment"
}
