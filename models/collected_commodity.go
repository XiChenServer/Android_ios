package models

import "gorm.io/gorm"

type CollectedCommodity struct {
	gorm.Model
	UserBasicID      uint
	CommodityBasicID uint
}

func (CollectedCommodity) TableName() string {
	return "collected_commodity"
}
