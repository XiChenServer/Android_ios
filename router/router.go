package router

import (
	"Android_ios/middleware"
	"Android_ios/servers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.POST("/send_phone_code", servers.BasicServer{}.SendPhoneCode)

	r.POST("/user/register/phone", servers.BasicOperateUser{}.UserRegisterByPhone)

	r.POST("/user/login/phone", servers.BasicOperateUser{}.UserLoginByPhoneCode)
	r.POST("/user/login/password", servers.BasicOperateUser{}.UserLoginByPassword)
	r.POST("/user/login/phone_and_password", servers.BasicOperateUser{}.UserLoginByPhoneAndPassword)

	user := r.Group("/user", middleware.AuthMiddleware())
	{
		user.POST("/uploads/avatar", servers.BasicOperateUser{}.UserUploadsAvatar)
		user.GET("/get/avatar", servers.BasicOperateUser{}.UserGetAvatar)
		user.GET("/get/info", servers.BasicOperateUser{}.UserGetInfo)
	}

	return r
}

// CORSMiddleware 处理跨域请求的中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的请求来源，你可以根据实际需要进行调整
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		// 允许客户端发送cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理预检请求（OPTIONS请求）
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 执行后续的处理函数
		c.Next()
	}
}
