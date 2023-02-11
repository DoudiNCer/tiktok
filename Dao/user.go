package Dao

import (
	"time"
)

// User User实体类
type User struct {
	Id         int       `gorm:"column:id"`          //用户id
	Name       string    `gorm:"column:name"`        // 用户名称
	Password   string    `gorm:"column:password"`    //用户密码
	CreateTime time.Time `gorm:"column:create_time"` //用户创建时间
}

func (ui User) TableName() string {
	return "User"
}

// UserResp 响应结构体
type UserResp struct {
	Id            uint   `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}
