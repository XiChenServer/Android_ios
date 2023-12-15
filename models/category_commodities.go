package models

import "gorm.io/gorm"

type CategoryCommodities struct {
	gorm.Model
	KindIdentity      string `gorm:"column:kind_identity;type:varchar(36);" json:"kind_identity"`
	CommodityIdentity string `gorm:"column:commodity_identity;type:varchar(36);" json:"commodity_identity"`
}

func (CategoryCommodities) TableName() string {
	return "category_commodities"
}
