package model

import "time"

type Comment struct {
	Id          int64 `gorm:"primaryKey"`
	CreatorId   int64
	Text        string
	VideoId     int64
	IsDeleted   bool
	UpdateTime  time.Time
	CreatedTime time.Time
}

func (comment *Comment) TableName() string {
	return "comment"
}
