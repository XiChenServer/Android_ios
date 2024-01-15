package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserChatBasicServer struct {
}

// CreateUserChat
// @Summary 创建用户聊天账号信息
// @Description 从现有用户基本信息中创建用户聊天账号信息
// @Accept json
// @Produce json
// @Tags 用户聊天基础
// @Success 200 {string} json{"code": 200, "msg": "用户聊天账号信息创建成功"}
// @Failure 500 {string} json{"code": 500, "msg": "服务器内部错误"}
// @Router /chat/basic/create [post]
func (UserChatBasicServer) CreateUserChat(c *gin.Context) {
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
		"msg":  "用户聊天账号信息创建成功",
	})
}
