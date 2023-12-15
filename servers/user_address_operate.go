package servers

import (
	"Android_ios/models"
	"Android_ios/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressEntry struct {
	Street   string `json:"street"`    // 街道
	City     string `json:"city"`      // 城市
	Country  string `json:"country"`   // 国家
	Province string `json:"province"`  // 省份
	Contact  string `json:"contact"`   // 联系人及其信息
	PostCode string `json:"post_code"` // 邮件地址
	Identity string `json:"identity"`  // 唯一标识
}

// UserUploadAddress 向用户上传地址信息
// @Summary 向用户上传地址信息
// @Description 通过解析用户信息和JSON请求体，将地址信息上传到用户信息中。
// @Tags 用户私有方法
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param address body AddressEntry true "要上传的地址信息,identity是不需要填写的，post_code可以写，也可以不写"
// @Success 200 {string} json {"code": 200, "msg": "地址添加成功"}
// @Failure 400 {string} json {"code": 400, "msg": "请求无效。服务器无法理解请求"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 409 {string} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/upload/address [post]
func (BasicOperateUser) UserUploadAddress(c *gin.Context) {
	// 解析用户信息
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	// 将 userClaim 转换为你的 UserClaims 结构体
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	// 使用 userClaims 中的信息获取文件名等等
	exists, existsuser, err := models.UserBasic{}.FindUserByAccount(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	if !exists {
		// 用户已经存在，你可以通过 existingUser 使用用户信息
		c.JSON(http.StatusConflict, gin.H{
			"code": "409",
			"msg":  "该账号没有被注册",
		})
		return
	}
	var address AddressEntry
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
		return
	}
	modelAddr := models.AddressEntry{
		Street:   address.Street,
		City:     address.City,
		Country:  address.Country,
		Province: address.Province,
		Contact:  address.Contact,
		PostCode: address.PostCode,
		Identity: pkg.GetAccountNumber(),
	}
	existsuser.Address = append(existsuser.Address, modelAddr)
	if err = existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "地址添加成功",
	})
	return
}

// UserDeleteAddress 删除用户地址
// @Summary 删除用户地址
// @Description 根据用户身份标识删除用户的地址信息。
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param identity formData string true "要删除的地址身份标识"
// @Success 200 {string} json {"code": 200, "msg": "地址删除成功"}
// @Failure 400 {string} json {"code": 400, "msg": "请求无效。服务器无法理解请求"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 403 {string} json {"code": 403, "msg": "禁止访问"}
// @Failure 404 {string} json {"code": 404, "msg": "未找到资源"}
// @Failure 409 {string} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/delete/address [delete]
func (BasicOperateUser) UserDeleteAddress(c *gin.Context) {
	identity := c.PostForm("identity")
	fmt.Println("identity:", identity)
	// 解析用户信息
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}

	// 将 userClaim 转换为你的 UserClaims 结构体
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	// 使用 userClaims 中的信息获取文件名等等
	exists, existsuser, err := models.UserBasic{}.FindUserByAccount(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	if !exists {
		// 用户已经存在，你可以通过 existingUser 使用用户信息
		c.JSON(http.StatusConflict, gin.H{
			"code": "409",
			"msg":  "该账号没有被注册",
		})
		return
	}

	flag := false
	address := existsuser.Address
	fmt.Println(identity)
	// 使用 range 遍历切片，i 是索引，v 是值
	for i, v := range address {
		fmt.Println(v.Identity)
		if i < 0 || i >= len(address) {
			// 处理错误或继续下一次迭代
			continue
		}
		if v.Identity == identity { // 使用 == 检查相等
			flag = true
			// 切片删除元素的正确方式
			existsuser.Address = append(address[:i], address[i+1:]...)
			break // 找到匹配的元素后，终止循环
		}
	}

	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  "没有该地址",
		})
		return
	}

	// 这里可能需要判断保存是否成功，如果不成功再返回错误
	if err := existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "地址删除成功",
	})
}
