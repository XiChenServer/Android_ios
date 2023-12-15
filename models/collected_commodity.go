package models

import "gorm.io/gorm"

type CollectedCommodity struct {
	gorm.Model
	CommodityIdentity string `gorm:"column:commodity_identity;type:varchar(36);" json:"CommodityIdentity"`
	CollectedIdentity string `gorm:"column:collected_identity;type:varchar(36);" json:"collected_identity"`
}

func (CollectedCommodity) TableName() string {
	return "collected_commodity"
}
