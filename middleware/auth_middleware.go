package middleware

import (
	"Android_ios/pkg"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 是用于验证 JWT 令牌的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段的值
		tokenString := c.GetHeader("Authorization")

		// 如果 Authorization 字段不存在或为空，则返回未授权错误
		if tokenString == "" {
			c.JSON(401, gin.H{
				"code": "401",
				"msg":  "未提供授权令牌"})
			c.Abort()
			return
		}

		// 去除 Bearer 前缀
		tokenString = tokenString[7:]

		// 解析 JWT 令牌
		userClaims, err := pkg.AnalyseToken(tokenString, c)
		if err != nil {
			c.JSON(401, gin.H{
				"code": "401",
				"msg":  "未提供授权令牌"})
			c.Abort()
			return
		}

		// 将用户声明信息存储到请求上下文中
		c.Set(pkg.UserClaimsContextKey, userClaims)

		// 继续处理请求
		c.Next()
	}
}
