package mysql

import (
	"errors"
	"github.com/DodiNCer/tiktok/biz/model"
	"github.com/DodiNCer/tiktok/biz/model/user_gorm"
	"github.com/DodiNCer/tiktok/biz/mw"
	"github.com/DodiNCer/tiktok/biz/util"
	"strconv"
	"time"
)

func QueryUserByUid(uid int64) (*model.User, error) {

	db := DB.Model(model.User{})

	db = db.Where("id = ?", uid)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}
	var res *model.User
	if err := db.First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// Register 注册用户
func Register(username, password string) (int64, error) {

	var tempUser model.User
	DB.Where("name = ?", username).First(&tempUser)
	if tempUser.Id > 0 {
		return 0, errors.New("用户名已存在")
	}
	tempUser.Name = username
	tempUser.Password = util.ScryptPassword(password)
	tempUser.CreateTime = time.Now()
	tempUser.PortraitPath = mw.MinioLinkPrefix + "head.jpg"
	tempUser.BackgroundPicturePath = mw.MinioLinkPrefix + "background.jpg"
	tempUser.Signature = "Welcome to tiktok"
	tx := DB.Begin()
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
	DB.Where("name = ?", username).First(&tempUser)
	if tempUser.Id <= 0 {
		return 0, errors.New("用户名不存在")
	}
	if tempUser.Password != util.ScryptPassword(password) {
		return 0, errors.New("密码错误")
	}
	return tempUser.Id, nil
}

// UserInfo 获取用户信息
func UserInfo(userId, Tid int64) (user_gorm.UserInfoResponse, error) {
	var user model.User
	var video model.Video
	var favorite model.Favorite
	var follower model.Follower
	var UserResp user_gorm.UserResp
	var userInfoResponse user_gorm.UserInfoResponse

	tx := DB.Begin()

	if err := tx.Where("id = ?", userId).First(&user).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("该用户不存在")
	}
	UserResp.Id = userId
	UserResp.Name = user.Name
	UserResp.Avatar = user.PortraitPath
	UserResp.BackgroundImage = user.BackgroundPicturePath
	UserResp.Signature = user.Signature

	var FollowerCount int64
	if err := tx.Where("user_uid=?", userId).Find(&follower).Count(&FollowerCount).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("粉丝总数统计异常")
	}
	UserResp.FollowerCount = FollowerCount

	var FollowCount int64
	if err := tx.Where("to_user_uid=?", userId).Find(&follower).Count(&FollowCount).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("关注总数统计异常")
	}
	UserResp.FollowCount = FollowCount

	tx.Where("to_user_uid = ?", Tid).First(&follower)
	if follower.Id <= 0 {
		UserResp.IsFollow = false
	} else {
		UserResp.IsFollow = true
	}
	//获赞数量
	totalFavorited, err := QueryNumOfFavoriteGotByUser(userId)
	if err != nil {
		return userInfoResponse, errors.New("获赞总数统计异常")
	}
	UserResp.TotalFavorited = strconv.FormatInt(totalFavorited, 10)

	//统计作品数
	var WorkCount int64
	if err := tx.Where("creator_id=?", userId).Find(&video).Count(&WorkCount).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("作品数统计异常")
	}
	UserResp.WorkCount = WorkCount

	//统计喜欢数
	var FavoriteCount int64
	if err := tx.Where("creator_id=?", userId).Find(&favorite).Count(&FavoriteCount).Error; err != nil {
		tx.Rollback()
		return userInfoResponse, errors.New("喜欢数统计异常")
	}
	UserResp.FavoriteCount = FavoriteCount

	userInfoResponse.UserResp = UserResp

	return userInfoResponse, tx.Commit().Error
}
