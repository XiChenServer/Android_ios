package models

import (
	"Android_ios/dao"
	"gorm.io/gorm"
)

type KindBasic struct {
	gorm.Model
	Name         string            `gorm:"column:name;type:varchar(36);" json:"name"`
	KindIdentity string            `gorm:"column:kind_identity;type:varchar(36);" json:"kind_identity"`
	ParentID     uint              `gorm:"column:parent_id;" json:"parent_id"`
	Parent       *KindBasic        `json:"parent"`
	Children     []*KindBasic      `gorm:"foreignKey:ParentID" json:"children"`
	Commodities  []*CommodityBasic `gorm:"many2many:kind_commodity_relations" json:"commodities"`
}

func (KindBasic) TableName() string {
	return "kind_basic"
}
func (KindBasic) GetKindBasicLink() (*KindBasic, error) {
	var kindBasic KindBasic
	err := dao.DB.First(&kindBasic).Error
	return &kindBasic, err
}

func (KindBasic) GetKindCommodityLink() (*KindBasic, error) {
	var kindBasic KindBasic
	err := dao.DB.Preload("Commodities").First(&kindBasic).Error
	return &kindBasic, err
}

// isValidType 验证商品类型是否有效
func IsValidType(typeName string) bool {
	// 查询数据库，检查商品类型是否存在于 KindBasic 表中
	var kind KindBasic
	if err := dao.DB.Where("name = ?", typeName).First(&kind).Error; err != nil {
		// 处理错误，例如类型不存在
		return false
	}

	// 如果找到相应的 KindBasic 记录，表示类型有效
	return true
}

// 在 KindBasic 模型中添加 FindKindByKindIdentity 方法
func (kb *KindBasic) FindKindByKindIdentity(name string) (*KindBasic, error) {
	var kind KindBasic
	if err := dao.DB.Where("name = ?", name).First(&kind).Error; err != nil {
		return nil, err
	}
	return &kind, nil
}
