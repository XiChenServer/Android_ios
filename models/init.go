package models

import "Android_ios/dao"

func Init() {

	// CategoryCommodities

	dao.DB.AutoMigrate(&KindCommodityRelation{})
	// LikedCommodity and CollectedCommodity
	dao.DB.AutoMigrate(&LikedCommodity{})
	dao.DB.AutoMigrate(&CollectedCommodity{})
	// KindBasic
	dao.DB.AutoMigrate(&KindBasic{})
	// CommodityBasic
	dao.DB.AutoMigrate(&CommodityBasic{})

	// UserBasic
	dao.DB.AutoMigrate(&UserBasic{})

	// UserChatBasic
	dao.DB.AutoMigrate(&UserChatBasic{})
	dao.DB.AutoMigrate(&Message{})
	dao.DB.AutoMigrate(&Contact{})
}
