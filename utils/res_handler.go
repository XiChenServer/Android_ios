package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*<<===================================admin===========================>>*/

type Admin struct {
}

func (Admin) AdminServer(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"code": code, "msg": msg})
}
func (Admin) AdminServerAndData(c *gin.Context, code int, msg string, data interface{}) {
	if code == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取分类信息成功", "data": data})
	}
	c.JSON(code, gin.H{"code": code, "msg": msg})
}

/*<<===================================basicServer===========================>>*/

type BasicServer struct {
}

// 发送验证码的响应
func (BasicServer) SendCodeServer(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"code": code, "msg": msg})
}
func (BasicServer) AdminServerAndData(c *gin.Context, code int, msg string, data interface{}) {
	if code == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取分类信息成功", "data": data})
	}
	c.JSON(code, gin.H{"code": code, "msg": msg})
}
