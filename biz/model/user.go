package model

import "time"

type User struct {
	Id         int64     `json:"id" column:"id"`
	Name       string    `json:"name" column:"name"`
	Password   string    `json:"password" column:"password"`
	CreateTime time.Time `json:"create_time" column:"create_time"`
}

func (u *User) TableName() string {
	return "user"
}
