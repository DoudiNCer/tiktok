package model

import (
	"time"
)

type Favorite struct {
	Id         int64 `gorm:"primaryKey;autoIncrement"`
	CreatorId  int64
	VideoId    int64
	IsDeleted  bool
	UpdateTime time.Time
	CreateTime time.Time
}

func (f *Favorite) TableName() string {
	return "favorite"
}
