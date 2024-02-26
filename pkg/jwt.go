package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// UserClaimsContextKey 是用于在上下文中存储用户声明的键
const UserClaimsContextKey = "UserClaims"

type UserClaims struct {
	UserIdentity string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	Account      string `json:"account"`
	PhoneNumber  string `json:"phone_number"`
	jwt.StandardClaims
}

// 获取密钥，如果环境变量中没有设置，则使用默认值
var myKey = []byte(getSecretKey())

func getSecretKey() string {
	// 从环境变量中获取密钥
	secretKey := os.Getenv("MY_SECRET_KEY")

	// 如果环境变量不存在或为空，使用默认值
	if secretKey == "" {
		secretKey = "default_secret_key"
	}
	fmt.Println("fsdfdf", secretKey)
	return secretKey
}

// GenerateToken 生成带有用户声明的JWT令牌
func GenerateToken(userIdentity, account, phoneNumber string) (string, error) {
	userClaims := &UserClaims{
		UserIdentity: userIdentity,
		Account:      account,
		PhoneNumber:  phoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", fmt.Errorf("签署令牌失败：%v", err)
	}

	return tokenString, nil
}

// AnalyseToken 解析JWT令牌并返回用户声明
func AnalyseToken(tokenString string, c *gin.Context) (*UserClaims, error) {
	userClaims := new(UserClaims)
	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("令牌格式错误")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("令牌已过期或尚未生效")
			} else {
				return nil, fmt.Errorf("令牌验证错误：%v", err)
			}
		}
		return nil, fmt.Errorf("解析令牌失败：%v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("令牌无效")
	}

	// 将解析后的用户信息存储到请求上下文中
	c.Set(UserClaimsContextKey, userClaims)

	return userClaims, nil
}
