package servers

import (
	"Android_ios/dao"
	"Android_ios/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Basic struct {
	Phone string `json:"phone"`
}

type BasicServer struct {
}

// SendPhoneCode 发送手机验证码
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
	err := pkg.SendPhoneCode(user.Phone, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	err = dao.RDB.Set(c, user.Phone, code, 300*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "验证码已经成功发送",
	})
	return
}
