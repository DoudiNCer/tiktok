package model

import "time"

type Video struct {
	Id         int64 `gorm:"primaryKey;autoIncrement"`
	Title      string
	Path       string
	CreatorId  int64
	CreateTime time.Time
	CoverPath  string
	IsDeleted  bool
	UpdateTime time.Time
}

func (video *Video) TableName() string {
	return "video"
}
