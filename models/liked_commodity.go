package models

import "gorm.io/gorm"

type LikedCommodity struct {
	gorm.Model
	UserBasicID      uint
	CommodityBasicID uint
}

func (LikedCommodity) TableName() string {
	return "liked_commodity"
}
