package user_gorm

// BaseResponse 基本响应
type BaseResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// UserResponse 用户注册和登录的共用响应
type UserResponse struct {
	BaseResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	BaseResponse
	UserResp `json:"user"` //用户信息
}

// UserResp 响应结构体
type UserResp struct {
	Id            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}
