package models

import "gorm.io/gorm"

type LikedCommodity struct {
	gorm.Model
	CommodityIdentity string `gorm:"column:commodity_identity;type:varchar(36);" json:"commodity_identity"`
	LikedIdentity     string `gorm:"column:liked_identity;type:varchar(36);" json:"liked_identity"`
}

func (LikedCommodity) TableName() string {
	return "liked_commodity"
}
