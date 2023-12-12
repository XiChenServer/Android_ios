package models

import "Android_ios/dao"

func Init() {
	dao.DB.AutoMigrate(&UserBasic{})
}
