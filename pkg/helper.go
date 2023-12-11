package pkg

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

// 加密
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// 随机获取十一位的账号
func GetAccountNumber() string {
	rand.Seed(time.Now().UnixNano())

	// 生成11位数字，转换为字符串
	randomNumber := rand.Intn(100000000000)
	return fmt.Sprintf("%011d", randomNumber)
}
func GetRandCode() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成六位随机验证码
	code := rand.Intn(900000) + 100000

	return fmt.Sprintf("%06d", code)
}

type UserClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
	IsAdmin int    `json:"is_admin"`
	Account string `json:"identity"`
	Phone   string `json:"phone"`
}

var myKey = []byte("new_system_server")

// 生成token
func GenerateToken(account, phone, name string, is_admin int) (string, error) {
	UserClaims := &UserClaims{
		IsAdmin:        is_admin,
		Name:           name,
		Phone:          phone,
		Account:        account,
		StandardClaims: jwt.StandardClaims{},
	}
	var myKey = []byte("new_system_server")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	UserClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, UserClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("签名验证失败")
		}
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("令牌无效")
	}
	return UserClaims, nil
}

// 生成随机昵称的函数
func GenerateRandomCreativeNickname() string {
	rand.Seed(time.Now().UnixNano())

	// 常见中文姓氏和名字
	chineseLastNames := []string{"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴"}
	chineseFirstNames := []string{"明", "丽", "强", "美", "刚", "秀", "伟", "娟", "勇", "芳"}

	// 常见英文名字
	englishFirstNames := []string{"Tom", "Jerry", "Alice", "Bob", "Lily", "John", "Emma", "David", "Sophie", "Michael"}

	// 创意组合
	creativeCombos := []string{"可爱的小", "搞怪的", "热情的", "冷静的", "幸运的", "聪明的", "快乐的", "迷人的"}

	// 生成随机昵称
	firstName := ""
	if rand.Intn(2) == 0 {
		// 使用中文名字
		firstName = chineseFirstNames[rand.Intn(len(chineseFirstNames))]
	} else {
		// 使用英文名字
		firstName = englishFirstNames[rand.Intn(len(englishFirstNames))]
	}

	lastName := chineseLastNames[rand.Intn(len(chineseLastNames))]
	creativeCombo := creativeCombos[rand.Intn(len(creativeCombos))]

	// 返回生成的随机昵称
	return fmt.Sprintf("%s%s%s", creativeCombo, lastName, firstName)
}
