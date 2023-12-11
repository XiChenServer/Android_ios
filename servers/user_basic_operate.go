package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressEntry struct {
	Street   string `json:"street"`   // 街道
	City     string `json:"city"`     // 城市
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
}

type User struct {
	Avatar           string         `json:"avatar"`            // 头像链接
	UserIdentity     string         `json:"user_identity"`     // 用户身份信息
	NickName         string         `json:"nickname"`          // 昵称
	Account          string         `json:"account"`           // 账号
	Password         string         `json:"password"`          // 密码
	PhoneNumber      string         `json:"phone_number"`      // 电话号码
	Email            string         `json:"email"`             // 电子邮件
	WechatNumber     string         `json:"wechat_number"`     // 微信号
	Address          []AddressEntry `json:"address"`           // 地址列表
	Score            int            `json:"score"`             // 评分
	VerificationCode string         `json:"verification_code"` // 验证码
}
type BasicOperateUser struct {
}

func (BasicOperateUser) UserRegisterByPhone(c *gin.Context) {
	var user User
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 检查手机号是否已经注册
	exists, _, err := models.UserBasic{}.FindUserByPhone(user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询错误"})
		return
	}
	if exists {
		// 用户已经存在，你可以通过 existingUser 使用用户信息
		c.JSON(http.StatusConflict, gin.H{
			"code": "409",
			"msg":  "该手机号已经注册",
		})
		return
	}

	code, err := dao.RDB.Get(c, user.PhoneNumber).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	if code != user.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "验证码错误",
		})
		return
	}
	var account string
	for {
		account = pkg.GetAccountNumber()
		exists, _, err = models.UserBasic{}.FindUserByAccount(account)
		if !exists && err == nil {
			break
		}
	}
	nickname := pkg.GenerateRandomCreativeNickname()
	dao.DB.Create(&models.UserBasic{
		Avatar:       "",
		UserIdentity: "",
		NickName:     nickname,
		Account:      account,
		Password:     pkg.GetMd5(user.Password),
		PhoneNumber:  user.PhoneNumber,
		Email:        "",
		WechatNumber: "",
		Address:      nil,
		Score:        0,
	})
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "注册成功",
		"data": map[string]interface{}{
			"nickname": "您的昵称是: " + nickname,
			"account":  "您的账号是: " + account,
			"token":    "",
		},
	})
	return
}
