package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id         int64  `json:"id" column:"id"`
	Name       string `json:"name" column:"name"`
	Password   string `json:"password" column:"password"`
	CreateTime string `json:"create_time" column:"create_time"`
}
