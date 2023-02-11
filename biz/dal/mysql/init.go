package mysql

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	if err != nil {
		panic(err)
	}
}
