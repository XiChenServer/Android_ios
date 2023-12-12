package models

import (
	"Android_ios/dao"
	"errors"
	"gorm.io/gorm"
	"log"
)

type AddressEntry struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Province string `json:"province"`
	Contact  string `json:"contact"`
	PostCode string `json:"post_code"`
}
type UserBasic struct {
	Avatar       string         `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	UserIdentity string         `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	NickName     string         `gorm:"column:nickname;type:varchar(24);" json:"nickname"`
	Account      string         `gorm:"column:account;type:varchar(11);" json:"account"`
	Password     string         `gorm:"column:password;type:varchar(255);" json:"password"`
	PhoneNumber  string         `gorm:"column:phone_number;type:varchar(16);" json:"phone_number"`
	Email        string         `gorm:"column:email;type:varchar(24);" json:"email"`
	WechatNumber string         `gorm:"column:wechat_number;type:varchar(24);" json:"wechat_number"`
	Address      []AddressEntry `gorm:"column:address;type:json;" json:"address"`
	Score        int            `gorm:"column:score;type:int;" json:"score"`
	Name         string         `gorm:"column:name;type:varchar(24);" json:"name"`
}

func (UserBasic) Table() string {
	return "user_basic"
}

func (UserBasic) SaveUser(user *UserBasic) error {
	// 使用 GORM 连接数据库（这里使用 dao.DB，确保在你的代码中初始化了数据库连接）
	db := dao.DB

	// 执行保存操作
	result := db.Model(&UserBasic{}).Where("user_identity = ?", user.UserIdentity).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (UserBasic) FindUserByPhone(phone string) (bool, *UserBasic, error) {
	var user UserBasic
	result := dao.DB.Select("avatar, user_identity, nickname, account, phone_number, email, wechat_number, address, score").
		Where("phone_number = ?", phone).
		First(&user).Statement
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found for phone: %s", phone)
			return false, nil, nil
		}
		log.Printf("Error executing query: %v", result.Error)
		return false, nil, result.Error
	}

	return true, &user, nil
}

func (UserBasic) FindUserByAccount(account string) (bool, *UserBasic, error) {
	var user UserBasic
	result := dao.DB.Select("avatar, user_identity, nickname, account, phone_number, email, wechat_number, address, score").
		Where("account = ?", account).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 未找到记录
			return false, nil, nil
		}
		// 查询时发生错误
		return false, nil, result.Error
	}

	// 找到记录
	return true, &user, nil
}

func (UserBasic) FindUserByAccountAndPassword(account string) (bool, *UserBasic, error) {
	var user UserBasic
	result := dao.DB.Where("account = ?", account).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 未找到记录
			return false, nil, nil
		}
		// 查询时发生错误
		return false, nil, result.Error
	}

	// 找到记录
	return true, &user, nil
}
func (UserBasic) FindUserByPhoneAndPassword(phone string) (bool, *UserBasic, error) {
	var user UserBasic
	result := dao.DB.Where("phone_number = ?", phone).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 未找到记录
			return false, nil, nil
		}
		// 查询时发生错误
		return false, nil, result.Error
	}

	// 找到记录
	return true, &user, nil
}
