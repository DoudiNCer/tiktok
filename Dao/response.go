package Dao

// BaseResponse 基本响应
type BaseResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// UserResponse 用户注册和登录的共用响应
type UserResponse struct {
	BaseResponse
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}
