package models

import "gorm.io/gorm"

// 中间表模型
type KindCommodityRelation struct {
	gorm.Model
	KindBasicID      uint
	CommodityBasicID uint
}

func (KindCommodityRelation) TableName() string {
	return "kind_commodity_relations"
}
