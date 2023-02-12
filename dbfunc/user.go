package dbfunc

import (
	"errors"
	"github.com/DodiNCer/tiktok/biz/dal/sql"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/DodiNCer/tiktok/util"

	"time"
)

// Register 注册用户
func Register(username, password string) (int64, error) {

	//dsn := "root:asd020118@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	return -1, err
	//}

	var tempUser model.User
	sql.DB.Where("name = ?", username).First(&tempUser)
	if tempUser.Id > 0 {
		return 0, errors.New("用户名已存在")
	}
	tempUser.Name = username
	tempUser.Password = util.ScryptPassword(password)
	tempUser.CreateTime = time.Now()

	tx := sql.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}
	err := tx.Create(&tempUser).Error
	if err != nil {
		tx.Rollback()
		return 0, errors.New("注册失败")
	}
	err = tx.Where("name = ?", username).First(&tempUser).Error
	if err != nil || tempUser.Id <= 0 {
		tx.Rollback()
		return 0, err
	}

	return tempUser.Id, tx.Commit().Error
}

// Login 用户登录
func Login(username, password string) (int64, error) {
	var tempUser model.User
	sql.DB.Where("name = ?", username).First(&tempUser)
	if tempUser.Id <= 0 {
		return 0, errors.New("用户名不存在")
	}
	if tempUser.Password != util.ScryptPassword(password) {
		return 0, errors.New("密码错误")
	}
	return tempUser.Id, nil
}

// UserInfo 获取用户信息
func UserInfo(userId, Tid int64) (model.UserInfoResponse, error) {
	var user model.User
	var follower model.Follower
	var UserResp model.UserResp
	var userInfoResponse model.UserInfoResponse

	tx := sql.DB.Begin()

	if err := tx.Where("id = ?", userId).First(&user).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("该用户不存在")
	}
	UserResp.Id = userId
	UserResp.Name = user.Name

	var FollowCount int64
	if err := tx.Where("user_uid=?", userId).Find(&follower).Count(&FollowCount).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("粉丝总数统计异常")
	}
	UserResp.FollowCount = FollowCount

	var FollowerCount int64
	if err := tx.Where("to_user_uid=?", userId).Find(&follower).Count(&FollowerCount).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("关注总数统计异常")
	}
	UserResp.FollowerCount = FollowerCount

	tx.Where("to_user_uid = ?", Tid).First(&follower)
	if follower.Id <= 0 {
		UserResp.IsFollow = false
	} else {
		UserResp.IsFollow = true
	}
	userInfoResponse.UserResp = UserResp

	return userInfoResponse, tx.Commit().Error
}
