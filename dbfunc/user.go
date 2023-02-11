package dbfunc

import (
	"errors"
	"github.com/DodiNCer/tiktok/Dao"
	"github.com/DodiNCer/tiktok/biz/dal/sql"
	"github.com/DodiNCer/tiktok/util"

	"time"
)

// Register 注册用户
func Register(username, password string) (int, error) {

	//dsn := "root:asd020118@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	return -1, err
	//}

	var tempUser Dao.User
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
