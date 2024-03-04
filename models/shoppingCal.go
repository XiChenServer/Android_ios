package models

import "gorm.io/gorm"

type ShoppingCar struct {
	gorm.Model
	Name       string  `gorm:"column:name;type:varchar(255);" json:"name"`
	Image      string  `gorm:"column:image;type:varchar(255);" json:"image"`
	Price      float64 `gorm:"column:price;" json:"price"`
	ProductId  uint    `gorm:"column:product_id;" json:"product_id"`
	UserId     string  `gorm:"column:user_id;type:varchar(255);" json:"user_id"`
	ProductNum uint    `gorm:"column:product_num;" json:"product_num"`
}

func (ShoppingCar) TableName() string {
	return "shopping_car"
}
