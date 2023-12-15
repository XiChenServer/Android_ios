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

	//

}
