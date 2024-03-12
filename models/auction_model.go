package models

import (
	"gorm.io/gorm"
	"time"
)

// 定义竞拍记录结构体

type BidRecordServer struct {
}

// BidRecord 结构体用于表示每次竞拍的信息
type BidRecord struct {
	gorm.Model
	ProductId   uint      `gorm:"product_id:commodity_identity;" json:"product_id"`
	UserID      string    `gorm:"column:user_id;type:varchar(36);" json:"user_id"`
	UserAccount string    `gorm:"column:user_account;type:varchar(11);" json:"user_account"`
	Price       float64   `gorm:"column:price;type:decimal(10,2)" json:"price"`
	BidTime     time.Time `gorm:"column:bid_time" json:"bid_time"`
}

func (BidRecord) TableName() string {
	return "bid_record"
}

// UserBidRecord 结构体用于表示当前竞拍信息
type UserBidRecord struct {
	ProductId       uint      `gorm:"product_id:commodity_identity;" json:"product_id"`
	UserID          string    `gorm:"column:user_id;type:varchar(36);" json:"user_id"`
	CurrentBidPrice float64   `json:"current_bid_price"`
	CurrentBidTime  time.Time `json:"current_bid_time"`
}

func (UserBidRecord) TableName() string {
	return "user_bid_record"
}
