package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Basic struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type BasicServer struct {
}

// SendPhoneCode 向指定手机号发送验证码。
// @Tags 公共方法
// @Summary 发送手机验证码
// @Description 向指定手机号发送验证码。
// @Accept json
// @Produce json
// @Param phone body Basic true "接收验证码的手机号"
// @Success 201 {string} json {"code":201,"msg":"验证码已经成功发送"}
// @Failure 400 {string} json {"code":400,"msg":"请求无效。服务器无法理解请求"}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /send_phone_code [post]
func (BasicServer) SendPhoneCode(c *gin.Context) {
	var user Basic
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
		return
	}

	code := pkg.GetRandCode()
	err = dao.RDB.Set(c, user.PhoneNumber, code, 300*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误,未能存储到redis",
		})
		return
	}

	err := pkg.SendPhoneCode(user.PhoneNumber, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误，阿里云发送失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码已经成功发送",
	})
	return
}

// SendEmailCode 向指定手机号发送验证码。
// @Tags 公共方法
// @Summary 发送邮箱验证码
// @Description Send an email verification code to the specified email address
// @Produce json
// @Param email formData string true "Email address to send the verification code to"
// @Success 200 {string} json {"code":200,"msg":"验证码已经成功发送"}
// @Failure 400 {string} json {"code":400,"msg":"请求参数错误"}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /send_email_code [post]
func (BasicServer) SendEmailCode(c *gin.Context) {
	email := c.PostForm("email")
	code := pkg.GetRandCode()
	err := pkg.SendEmailVerificationCode(email, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	err = dao.RDB.Set(c, email, code, 300*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码已经成功发送",
	})
	return
}

func (BasicServer) RenewUserChat(c *gin.Context) {
	var users []models.UserBasic
	err := dao.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	for _, user := range users {
		fmt.Println(user)
		err := models.UserChatBasic{}.CreateUser(user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务器内部错误",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户聊天信息更新成功",
	})
}
