package models

import (
	"Android_ios/dao"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type MediaBasic struct {
	Image string `json:"image"`
}

type JSONMedia []MediaBasic

// JSONMedia 定义...
func (jm *JSONMedia) Scan(value interface{}) error {
	if value == nil {
		*jm = nil
		return nil
	}
	err := json.Unmarshal(value.([]byte), jm)
	if err != nil {
		fmt.Printf("解组 JSON 时出错：%v\n", err)
	}
	return err
}

func (jm JSONMedia) Value() (driver.Value, error) {
	if jm == nil {
		return nil, nil
	}
	return json.Marshal(jm)
}

// CommodityBasic 表示商品的基本信息。
type CommodityBasic struct {
	gorm.Model
	CommodityIdentity string `gorm:"column:commodity_identity;type:varchar(36);index" json:"commodity_identity"` // CommodityIdentity 是商品的唯一标识符。
	//UserInfo          *UserBasic  `gorm:"foreignKey:UserIdentity;references:CommodityIdentity"`
	Title        string      `gorm:"column:title;type:varchar(36);" json:"title"`     // Title 是商品的标题或名称。
	Number       int         `gorm:"column:number" json:"number"`                     // Number 表示商品的数量或数量。
	Information  string      `gorm:"column:information;type:text" json:"information"` // Information 提供有关商品的额外详细信息。
	Price        float64     `gorm:"column:price;type:decimal(10,2)" json:"price"`    // Price 表示商品的价格。
	SoldStatus   int         `gorm:"column:sold_status" json:"sold_status"`           // SoldStatus 表示商品是否已售出的状态。1:正在出售中。2：正在拍卖中。: 提交订单，正在发货 3：已经卖完
	Media        JSONMedia   `gorm:"column:media;type:json" json:"media"`             // Media 包含与商品相关的嵌套媒体信息。
	IsAuction    int         `gorm:"column:is_auction" json:"is_auction"`             // 0是不拍卖，1是拍卖                                   // IsAuction 表示商品是否属于拍卖。
	Address      JSONAddress `gorm:"column:address;type:json" json:"address"`         // Address 包含与商品相关的嵌套地址信息。
	LikeCount    int         `gorm:"column:like_count" json:"like_count"`             // LikeCount 表示商品收到的点赞数。
	CollectCount int         `gorm:"column:collect_count" json:"collect_count"`       // CollectCount 表示商品被收藏的次数。

	Categories []*KindBasic `gorm:"many2many:kind_commodity_relations" json:"categories"` // Categories 表示与商品相关的类别。
	LikedUsers []*KindBasic `gorm:"many2many:liked_commodity" json:"liked_users"`
}

func (CommodityBasic) TableName() string {
	return "commodity_basic"
}
func (CommodityBasic) CreateCommodity(commodity *CommodityBasic) error {
	// 使用 GORM 连接数据库（这里使用 dao.DB，确保在你的代码中初始化了数据库连接）
	db := dao.DB

	// 执行创建操作
	result := db.Create(commodity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindCommodityByID 根据商品ID查找商品
func (cb *CommodityBasic) FindCommodityByID(commodityID string) (*CommodityBasic, error) {
	var commodity CommodityBasic
	if err := dao.DB.Where("commodity_identity = ?", commodityID).Preload("Categories").Preload("Media").Preload("Address").First(&commodity).Error; err != nil {
		return nil, err
	}
	return &commodity, nil
}

// UpdateCommodity 更新商品信息
func (c *CommodityBasic) UpdateCommodity(updates CommodityBasic) error {
	// 使用 Model 方法指定要更新的记录
	if err := dao.DB.Model(c).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}
