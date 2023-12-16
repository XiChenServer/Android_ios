package models

type Admin struct {
	Account  string `gorm:"column:account;type:varchar(11);" json:"account"`
	Password string `gorm:"column:password;type:varchar(255);" json:"password"`
}
