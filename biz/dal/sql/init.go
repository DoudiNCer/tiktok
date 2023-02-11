package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDb gorm初始化函数
func InitDb() error {
	var err error
	dsn := "tiktok:tiktok@tcp(127.0.0.1:13306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
