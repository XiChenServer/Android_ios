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
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码已经成功发送",
	})
	return
}
