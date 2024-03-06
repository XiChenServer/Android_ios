package models

import (
	"Android_ios/dao"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderIdentity   string         `gorm:"column:order_identity;type:varchar(36);" json:"order_identity"`
	BuyerIdentity   string         `gorm:"column:buyer_identity;type:varchar(36);" json:"buyer_identity"`
	SellerIdentity  string         `gorm:"column:seller_identity;type:varchar(36);" json:"seller_identity"`
	ProductIdentity uint           `gorm:"column:product_id;type:uint;" json:"product_id"`
	Name            string         `gorm:"column:name;type:varchar(255);" json:"name"`
	Price           float64        `gorm:"column:price;type:decimal(10,2);" json:"price"`
	Quantity        int            `gorm:"column:quantity;type:int;" json:"quantity"`
	Msg             string         `gorm:"column:msg;type:text;" json:"msg"`
	Buyer           UserBasic      `gorm:"foreignKey:BuyerIdentity;references:UserIdentity"`
	Seller          UserBasic      `gorm:"foreignKey:SellerIdentity;references:UserIdentity"`
	Product         CommodityBasic `gorm:"foreignKey:ProductIdentity"`
	Status          uint           `gorm:"column:status;type:uint;" json:"status"` // 1：待支付 2：支付完成，开始发货 3：订单完成 4：订单取消
}

func (Order) TableName() string {
	return "order_basic"
}

func (Order) CreateOrder(order *Order) error {
	err := dao.DB.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}
func (Order) ModifyOrder(order Order) error {
	err := dao.DB.Updates(&order).Error
	if err != nil {
		return err
	}
	return nil
}
func (Order) DelOrder(order Order) error {
	err := dao.DB.Delete(&order).Error
	if err != nil {
		return err
	}
	return nil
}
