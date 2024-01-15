package models

import (
	"Android_ios/dao"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserChatBasic struct {
	gorm.Model
	UserIdentity  string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	NickName      string `gorm:"column:nickname;type:varchar(24);" json:"nickname"`
	Account       string `gorm:"column:account;type:varchar(11);" json:"account"`
	Avatar        string //头像
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (UserChatBasic) TableName() string {
	return "user_chat_basic"
}

func (UserChatBasic) GetUserList() []*UserChatBasic {
	data := make([]*UserChatBasic, 10)
	dao.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

/*func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	dao.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := dao.MD5Encode(str)
	dao.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}*/

func (UserChatBasic) FindUserByName(name string) UserChatBasic {
	user := UserChatBasic{}
	dao.DB.Where("name = ?", name).First(&user)
	return user
}
func (UserChatBasic) FindUserByPhone(phone string) *gorm.DB {
	user := UserChatBasic{}
	return dao.DB.Where("Phone = ?", phone).First(&user)
}
func (UserChatBasic) FindUserByEmail(email string) *gorm.DB {
	user := UserChatBasic{}
	return dao.DB.Where("email = ?", email).First(&user)
}
func (UserChatBasic) CreateUser(user UserBasic) *gorm.DB {
	return dao.DB.Create(&UserChatBasic{
		UserIdentity:  user.UserIdentity,
		NickName:      user.NickName,
		Account:       user.Account,
		Avatar:        user.Avatar,
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		LoginOutTime:  time.Now(), // 确保 login_out_time 字段使用零值，或者设置为一个合适的时间值

	})
}
func (UserChatBasic) DeleteUser(user UserBasic) *gorm.DB {
	return dao.DB.Delete(&user).Where("identity = ?", user.UserIdentity)
}

func (UserChatBasic) UpdateUser(user UserChatBasic) *gorm.DB {
	return dao.DB.Model(&user).Updates(UserChatBasic{
		UserIdentity: user.UserIdentity,
		NickName:     user.NickName,
		Account:      user.Account,
		Avatar:       user.Avatar,
	})
}

// 查找某个用户
func (UserChatBasic) FindByAccount(account string) UserChatBasic {
	user := UserChatBasic{}
	dao.DB.Where("account = ?", account).First(&user)
	return user
}
