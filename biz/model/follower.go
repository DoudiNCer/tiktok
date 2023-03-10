package model

import "time"

type Follower struct {
	Id         int64     `json:"id" column:"id"`
	ToUserUid  int64     `json:"to_user_uid" column:"to_user_uid"`
	UserUid    int64     `json:"user_uid" column:"user_uid"`
	CreateTime time.Time `json:"create_time" column:"create_time"`
	IsDeleted  bool      `json:"is_delete" column:"is_delete"`
	UpdateTime time.Time `json:"update_time" column:"update_time"`
}

func (f *Follower) TableName() string {
	return "follower"
}
