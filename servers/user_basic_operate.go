package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type AddressEntry struct {
	Street   string `json:"street"`    // 街道
	City     string `json:"city"`      // 城市
	Country  string `json:"country"`   // 国家
	Province string `json:"province"`  // 省份
	Contact  string `json:"contact"`   // 联系人及其信息
	PostCode string `json:"post_code"` // 邮件地址
}
type User struct {
	Avatar           string         `json:"avatar"`            // 头像链接
	UserIdentity     string         `json:"user_identity"`     // 用户身份信息
	NickName         string         `json:"nickname"`          // 昵称
	Account          string         `json:"account"`           // 账号
	Password         string         `json:"password"`          // 密码
	PhoneNumber      string         `json:"phone_number"`      // 电话号码
	Email            string         `json:"email"`             // 电子邮件
	WechatNumber     string         `json:"wechat_number"`     // 微信号
	Address          []AddressEntry `json:"address"`           // 地址列表
	Score            int            `json:"score"`             // 评分
	VerificationCode string         `json:"verification_code"` // 验证码
	Name             string         `json:"name"`
	FileBase64       string         `json:"file_base64"`
}
type BasicOperateUser struct {
}

// UserRegisterByPhone 向系统注册用户通过手机号。
// @Summary 用户注册
// @Description 用户通过手机号进行注册，如果手机号已存在或验证码错误将返回相应的错误信息。
// @Tags 用户方法
// @Accept json
// @Produce json
// @Param user body User true "用户信息"
// @Success 200 {string} json {"code": "200", "msg": "注册成功", "data": {"nickname": "用户昵称", "account": "用户账号": ""}}
// @Failure 400 {string} json {"code": "400", "msg": "请求无效。服务器无法理解请求"}
// @Failure 409 {string} json {"code": "409", "msg": "该手机号已经注册"}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /user/register/phone [post]
func (BasicOperateUser) UserRegisterByPhone(c *gin.Context) {
	var user User
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
		return
	}
	// 检查手机号是否已经注册
	exists, _, err := models.UserBasic{}.FindUserByPhone(user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	if exists {
		// 用户已经存在，你可以通过 existingUser 使用用户信息
		c.JSON(http.StatusConflict, gin.H{
			"code": "409",
			"msg":  "该手机号已经注册",
		})
		return
	}
	code, err := dao.RDB.Get(c, user.PhoneNumber).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "验证码没有发送",
		})
		return
	}
	if code != user.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "验证码错误",
		})
		return
	}
	var account string
	for {
		account = pkg.GetAccountNumber()
		exists, _, err = models.UserBasic{}.FindUserByAccount(account)
		if !exists && err == nil {
			break
		}
	}
	user_identity := pkg.GenerateUniqueID()
	nickname := pkg.GenerateRandomCreativeNickname()
	err = dao.DB.Create(&models.UserBasic{
		Avatar:       "",
		UserIdentity: user_identity,
		NickName:     nickname,
		Account:      account,
		Password:     pkg.GetHash(user.Password),
		PhoneNumber:  user.PhoneNumber,
		Email:        "",
		WechatNumber: "",
		Address:      nil,
		Score:        0,
	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "内部发生错误",
		})
		return
	}
	token, err := pkg.GenerateToken(user_identity, account, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "内部发生错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "注册成功",
		"data": map[string]interface{}{
			"nickname": "您的昵称是: " + nickname,
			"account":  "您的账号是: " + account,
			"token":    token,
		},
	})
	return
}

// UserLoginByPhoneCode 通过手机号和验证码进行用户登录。
// @Summary 用户登录（验证码）
// @Description 用户通过手机号和验证码进行登录，如果手机号未注册、验证码错误或其他错误将返回相应的错误信息。
// @Tags 用户方法
// @Accept json
// @Produce json
// @Param user body User true "用户信息"
// @Success 200 {string} json {"code": "200", "msg": "登录成功", "data": {"token": "用户token"}}
// @Failure 400 {string} json {"code": "400", "msg": "请求无效。服务器无法理解请求"}
// @Failure 409 {string} json {"code": "409", "msg": "该手机号没有被注册"}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /user/login/phone [post]
func (BasicOperateUser) UserLoginByPhoneCode(c *gin.Context) {
	var user User
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
	}
	// 检查手机号是否已经注册
	exists, existsuser, err := models.UserBasic{}.FindUserByPhone(user.PhoneNumber)

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
			"msg":  "该手机号没有被注册",
		})
		return
	}
	code, err := dao.RDB.Get(c, user.PhoneNumber).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "验证码没有发送",
		})
		return
	}
	if code != user.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "验证码错误",
		})
		return
	}
	token, err := pkg.GenerateToken(existsuser.UserIdentity, existsuser.Account, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "内部发生错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "登录成功",
		"data": map[string]interface{}{
			"token": token,
		},
	})
	return

}

// UserLoginByPassword 通过账号和密码进行用户登录。
// @Summary 用户登录（密码）
// @Description 用户通过账号和密码进行登录，如果账号未注册、密码错误或其他错误将返回相应的错误信息。
// @Tags 用户方法
// @Accept json
// @Produce json
// @Param user body User true "用户信息"
// @Success 200 {string} json {"code": "200", "msg": "登录成功", "data": {"token": "用户token"}}
// @Failure 400 {string} json {"code": "400", "msg": "请求无效。服务器无法理解请求"}
// @Failure 409 {string} json {"code": "409", "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /user/login/password [post]
func (BasicOperateUser) UserLoginByPassword(c *gin.Context) {
	var user User
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
	}
	// 检查账号是否已经注册
	exists, existsuser, err := models.UserBasic{}.FindUserByAccountAndPassword(user.Account)
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
	if existsuser.Password != pkg.GetHash(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "账号或密码错误",
		})
		return
	}

	token, err := pkg.GenerateToken(existsuser.UserIdentity, user.Account, existsuser.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "内部发生错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "登录成功",
		"data": map[string]interface{}{
			"token": token,
		},
	})
	return
}

// UserLoginByPhoneAndPassword 通过手机号进行用户登录。
// @Summary 用户登录（使用手机号和密码）
// @Description 用户通过手机号和密码进行登录
// @Tags 用户方法
// @Accept json
// @Produce json
// @Param user body User true "用户信息"
// @Success 200 {json} json "{"code": "200", "msg": "登录成功", "data": {"token": "jwt_token_here"}}"
// @Failure 400 {json} json "{"code": 400, "msg": "请求无效。服务器无法理解请求"}"
// @Failure 401 {json} json "{"code": 401, "msg": "该手机号没有被注册"}"
// @Failure 400 {json} json "{"code": 400, "msg": "手机号或密码错误"}"
// @Failure 500 {json} json "{"code": 500, "msg": "服务器内部错误"}"
// @Router /user/login/phone_and_password [post]
func (BasicOperateUser) UserLoginByPhoneAndPassword(c *gin.Context) {
	var user User
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
	}
	// 检查账号是否已经注册
	exists, existsuser, err := models.UserBasic{}.FindUserByPhoneAndPassword(user.PhoneNumber)
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
			"msg":  "该手机号没有被注册",
		})
		return
	}
	if existsuser.Password != pkg.GetHash(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "手机号或密码错误",
		})
		return
	}
	token, err := pkg.GenerateToken(existsuser.UserIdentity, existsuser.Account, existsuser.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "内部发生错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "登录成功",
		"data": map[string]interface{}{
			"token": token,
		},
	})
	return
}

// UserUploadsAvatar 处理用户上传头像的请求
// @Summary 上传用户头像
// @Description 处理用户上传头像的请求，需要提供有效的用户身份验证令牌。
// @Tags 用户私有方法
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param files formData file true "用户头像文件"
// @Success 200 {string} json {"code": 200, "msg": "文件上传成功"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 403 {string} json {"code": 403, "msg": "禁止访问"}
// @Failure 413 {string} json {"code": 413, "msg": "文件大小超出限制"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/uploads/avatar [post]
func (BasicOperateUser) UserUploadsAvatar(c *gin.Context) {
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	const MaxFileSize = 10 << 20
	// 将 userClaim 转换为你的 UserClaims 结构体
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	// 解析多部分表单数据
	err := c.Request.ParseMultipartForm(10 << 20) // 限制最大文件大小为 10MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	// 获取所有上传的文件
	form := c.Request.MultipartForm
	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	// 处理第一个上传的文件
	file := files[0]

	// 检查文件大小
	if file.Size > MaxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"code": http.StatusRequestEntityTooLarge, "msg": "文件大小超出限制"})
		return
	}
	// 生成唯一的文件名
	fileName := pkg.GenerateUniqueImageName(userClaims.UserIdentity, file.Filename)
	// 直接获取文件内容的 io.Reader
	fileReader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	defer fileReader.Close()
	// 在这里你可以使用 fileReader 进行进一步的操作
	// 例如，你可以将其传递给某个需要 io.Reader 的函数
	// DoSomethingWithReader(fileReader)
	// 保存文件
	// 检查账号是否已经注册
	exists, existsuser, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	if !exists {
		// 用户已经存在，你可以通过 existingUser 使用用户信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	existsuser.Avatar = fileName
	if err = existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	err = pkg.UploadAvatarFromForm(fileName, fileReader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "文件上传成功"})
}

// UserGetAvatar 从阿里云服务器获取用户头像
// @Summary 获取用户头像
// @Description 从阿里云服务器获取用户头像，需要提供有效的用户身份验证令牌。
// @Tags 用户私有方法
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}" default(123)
// @Success 200 {file} application/octet-stream "头像加载成功"
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 403 {string} json {"code": 403, "msg": "禁止访问"}
// @Failure 404 {string} json {"code": 404, "msg": "文件不存在"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/get/avatar [get]
func (BasicOperateUser) UserGetAvatar(c *gin.Context) {
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

	exists, existsuser, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
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

	fileName := existsuser.Avatar
	// 从阿里云 OSS 获取头像
	fileReader, err := pkg.DownAvatarFromOSS(fileName)
	fmt.Println(fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器内部错误2312" + err.Error(),
		})
		return
	}
	defer fileReader.Close() // 关闭文件

	// 读取文件内容到 []byte
	fileContent, err := ioutil.ReadAll(fileReader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "无法读取文件内容",
		})
		return
	}

	// 将自定义信息添加到 HTTP 头部
	c.Writer.Header().Set("Custom-Header", "头像加载成功")

	// 返回文件内容给客户端
	c.Data(http.StatusOK, "application/octet-stream", fileContent)
}

func (BasicOperateUser) UserModifyPassword(c *gin.Context) {
	var user User
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
	}
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

	exists, existsuser, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
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
	existsuser.Password = user.Password
}

// UserGetInfo 解析用户信息并获取用户信息的API
// @Summary 获取用户信息
// @Description 解析用户信息，验证用户身份，并返回用户信息。
// @Tags 用户私有方法
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}" default
// @Success 200 {json} json {"code": 200, "msg": "获取信息成功", "data": {用户信息}}
// @Failure 401 {json} json {"code": 401, "msg": "未授权"}
// @Failure 400 {json} json {"code": 400, "msg": "请求错误"}
// @Failure 404 {json} json {"code": 404, "msg": "用户不存在"}
// @Failure 409 {json} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {json} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/get/info [get]
func (BasicOperateUser) UserGetInfo(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取信息成功",
		"data": existsuser,
	})
	return
}
func (BasicOperateUser) UserModifyInfo(c *gin.Context) {
	var user User
	// 使用 ShouldBindJSON 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求无效。服务器无法理解请求",
		})
	}
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

	exists, existsuser, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
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
	if user.Password == "" || user.NickName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	existsuser.Password = user.Password
	existsuser.Name = user.Name
	existsuser.NickName = user.NickName
	existsuser.WechatNumber = user.WechatNumber
	if err = existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改信息成功",
	})
	return
}
