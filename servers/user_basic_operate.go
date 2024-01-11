package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

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
	dao.RDB.Del(c, user.PhoneNumber)
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
		Avatar:       "https://xichen-server/your-prefix/9682B802FF091A4746DCC98526E2FE8B.jpg",
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
	dao.RDB.Del(c, user.PhoneNumber)
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
	exists, existsuser, err := models.UserBasic{}.FindUserByPhoneAndPassword(user.PhoneNumber, pkg.GetHash(user.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	fmt.Println(existsuser)
	if !exists {
		// 用户已经存在，你可以通过 existingUser 使用用户信息
		c.JSON(http.StatusConflict, gin.H{
			"code": "409",
			"msg":  "该手机号没有被注册",
		})
		return
	}
	//if existsuser.Password != pkg.GetHash(user.Password) {
	//	fmt.Println("dfdf", user.Password)
	//	fmt.Println(pkg.GetHash(user.Password))
	//	fmt.Println(existsuser.Password)
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "手机号或密码错误",
	//	})
	//	return
	//}
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

	// 保存文件到本地
	localFilePath := "./picture/avatar/" + file.Filename
	fmt.Println("sdf", localFilePath)
	err = c.SaveUploadedFile(file, localFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 构建上传路径，直接使用文件名，确保包含目录信息
	objectKey := "picture/avatar/" + file.Filename
	// 上传本地文件到服务器
	//	objectKey := "/picture/avatar/" + file.Filename
	err = pkg.UploadAllFile(objectKey, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fileURL := fmt.Sprintf("http://127.0.0.1:13000/%s", objectKey)
	// 检查账号是否已经注册
	fmt.Println("URL", fileURL)
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
	existsuser.Avatar = fileURL
	if err = existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "文件上传成功"})
}

/*
// UserUploadLocal
// @Summary 上传用户头像
// @Tags 用户私有方法
// @Description 上传用户头像并更新用户信息
// @ID user-upload-local
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param files formData file true "用户头像文件"
// @Success 200 {string} json {"code": 200, "msg": "文件上传成功"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 403 {string} json {"code": 403, "msg": "禁止访问"}
// @Failure 413 {string} json {"code": 413, "msg": "文件大小超出限制"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/upload/local [post]
func (BasicOperateUser) UserUploadLocal(c *gin.Context) {
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
	suffix := ".png"
	ofilName := file.Filename
	tem := strings.Split(ofilName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstFile, _ := os.Create("./asset/upload/" + fileName)

	if err != nil {
		// 错误处理
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}
	defer dstFile.Close()

	srcFile, err := file.Open()
	if err != nil {
		// 错误处理
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	url := "./asset/upload/" + fileName
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
	existsuser.Avatar = url
	fmt.Println("fsddf", url)
	if err = existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "头像上传成功"})
}*/

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

// UserGetAvatarLocal
// @Summary 获取用户头像信息
// @Description 获取用户头像文件名等信息
// @ID user-get-avatar-local
// @Tag 用户私有信息
// @Produce json
// @Param account query string true "用户账号"
// @Success 200 {file} application/octet-stream "头像加载成功"
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 403 {string} json {"code": 403, "msg": "禁止访问"}
// @Failure 404 {string} json {"code": 404, "msg": "文件不存在"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/get/avatar/local [post]
func (BasicOperateUser) UserGetAvatarLocal(c *gin.Context) {
	account := c.Query("account")

	exists, existsuser, err := models.UserBasic{}.FindUserByAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器内部错误",
		})
		return
	}
	if !exists {
		// 用户已经存在，你可以通过 existingUser 使用用户信息
		c.JSON(http.StatusNotFound, gin.H{
			"code": "404",
			"msg":  "该账号没有被注册",
		})
		return
	}
	fileName := existsuser.Avatar
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "头像加载成功",
		"data": map[string]interface{}{
			"AvatarUrl": fileName,
		},
	})
	return
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

type UserModify struct {
	NickName         string `json:"nickname"`          // 昵称
	Password         string `json:"password"`          // 密码
	Email            string `json:"email"`             // 电子邮件
	WechatNumber     string `json:"wechat_number"`     // 微信号
	VerificationCode string `json:"verification_code"` // 验证码
	Name             string `json:"name"`
}

// UserModifyInfo 是一个Swagger文档化的函数，用于修改用户信息
// @Summary 修改用户信息
// @Description 根据提供的JSON请求体修改用户信息
// @Tag 用户私有方法
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}" default
// @Param user body UserModify true "需要修改的用户信息"
// @Param VerificationCode formData string false "验证码" maxlength(6) "验证码"
// @Success 200 {string} json {"code":200,"msg":"修改信息成功"}
// @Failure 400 {string} json {"code":400,"msg":"请求无效。服务器无法理解请求"}
// @Failure 401 {string} json {"code":401,"msg":"未授权"}
// @Failure 403 {string} json {"code":403,"msg":"请求错误"}
// @Failure 404 {string} json {"code":404,"msg":"该账号没有被注册"}
// @Failure 409 {string} json {"code":409,"msg":"请求错误"}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /user/modify/info [put]
func (BasicOperateUser) UserModifyInfo(c *gin.Context) {
	var user UserModify
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
	if user.Email != existsuser.Email {
		code, err := dao.RDB.Get(c, user.Email).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "验证码没有发送",
			})
			return
		}
		dao.RDB.Del(c, user.Email)
		if code != user.VerificationCode {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "验证码错误",
			})
			return
		}
		existsuser.Email = user.Email
	}
	existsuser.Password = user.Password
	existsuser.Name = user.Name
	existsuser.NickName = user.NickName
	existsuser.WechatNumber = user.WechatNumber
	existsuser.Email = user.Email
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

// UserChangesMobilePhoneNumber 更换用户手机号
// @Summary 更换用户手机号
// @Description 根据用户身份标识更换用户的手机号码。
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param phone_number formData string true "要更换的手机号码"
// @Param verification_code formData string true "验证码"
// @Success 200 {string} json {"code": 200, "msg": "手机号更换成功"}
// @Failure 400 {string} json {"code": 400, "msg": "请求无效。服务器无法理解请求"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 403 {string} json {"code": 403, "msg": "禁止访问"}
// @Failure 404 {string} json {"code": 404, "msg": "未找到资源"}
// @Failure 409 {string} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/changes/mobile/phone [post]
func (BasicOperateUser) UserChangesMobilePhoneNumber(c *gin.Context) {
	phone_number := c.PostForm("phone_number")
	verification_code := c.PostForm("verification_code")
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
	code, err := dao.RDB.Get(c, phone_number).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "验证码没有发送",
		})
		return
	}
	dao.RDB.Del(c, phone_number)
	if code != verification_code {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "验证码错误",
		})
		return
	}
	existsuser.PhoneNumber = phone_number
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
		"msg":  "手机号更换成功",
	})
	return
}

// UsersLikePro
// @Summary 点赞商品
// @Description 用户点赞商品
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param commodity_identity query string true "商品标识"
// @Success 200 {string} json {"code": 200, "msg": "点赞成功"}
// @Failure 401 {string} json {"code": 401, "message": "未授权"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 409 {string} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/like/product [post]
func (BasicOperateUser) UsersLikePro(c *gin.Context) {
	identity := c.Query("commodity_identity")
	id, err := strconv.Atoi(identity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "未授权"})
		return
	}

	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	_, existsuser, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	var pro models.CommodityBasic
	if err := dao.DB.Where("id = ?", uint(id)).First(&pro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	var product models.UserBasic
	if err := dao.DB.Preload("LikedCommodities").Where("user_identity = ?", userClaims.UserIdentity).
		First(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	// 检查用户是否已经点赞过该商品
	for _, likedProduct := range product.LikedCommodities {
		if likedProduct.ID == uint(id) {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "已经点赞过该商品"})
			return
		}
	}

	tx := dao.DB.Begin()

	pro.LikeCount++
	pro.UpdateCommodity(pro)

	product.LikedCommodities = append(product.LikedCommodities, &pro)
	existsuser.LikedCommodities = product.LikedCommodities
	if err := existsuser.SaveUser(existsuser); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "提交事务失败"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "提交事务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "点赞成功"})
}

// UserGetLikePro
// @Summary 获取用户点赞的商品数据
// @Description 获取用户点赞的商品数据
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} models.UserBasic "获取点赞的商品数据"
// @Failure 401 {string} json {"code": 401, "message": "未授权"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/get/like/pro [get]
func (BasicOperateUser) UserGetLikePro(c *gin.Context) {
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
	var product models.UserBasic
	if err := dao.DB.Preload("LikedCommodities").Where("user_identity = ?", userClaims.UserIdentity).
		First(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	var likedCommodities []*models.CommodityBasic
	likedCommodities = product.LikedCommodities
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取点赞的商品数据",
		"data": likedCommodities,
	})
}

// UsersUnlikePro
// @Summary 取消点赞商品
// @Description 用户取消点赞商品
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param commodity_identity query string true "商品标识"
// @Success 200 {string} json {"code": 200, "msg": "取消点赞成功"}
// @Failure 401 {string} json {"code": 401, "message": "未授权"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 409 {string} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/unlike/product [post]
func (u BasicOperateUser) UsersUnlikePro(c *gin.Context) {
	identity := c.Query("commodity_identity")
	commodityID, err := strconv.Atoi(identity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "商品标识解析错误"})
		return
	}

	// 解析用户信息
	userClaim, exist := c.Get(pkg.UserClaimsContextKey)
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "未授权"})
		return
	}

	// 将 userClaim 转换为你的 UserClaims 结构体
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	// 检查用户是否已注册
	userExists, _, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	if !userExists {
		c.JSON(http.StatusConflict, gin.H{"code": http.StatusConflict, "msg": "该账号没有被注册"})
		return
	}

	// 开启数据库事务
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询用户信息及点赞的商品
	var user models.UserBasic
	if err := tx.Preload("LikedCommodities").Where("user_identity = ?", userClaims.UserIdentity).
		First(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	if user.ID == 0 {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "未找到用户信息"})
		return
	}

	// 查询被取消点赞的商品
	var targetProduct models.CommodityBasic
	if err := tx.Where("id = ?", uint(commodityID)).First(&targetProduct).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "查询商品失败"})
		return
	}

	// 检查用户是否已经点赞过该商品
	var liked bool
	var likedIndex int
	for i, likedProduct := range user.LikedCommodities {
		if likedProduct.ID == uint(commodityID) {
			liked = true
			likedIndex = i
			break
		}
	}

	if !liked {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "未点赞过该商品"})
		return
	}

	// 从关联表中删除点赞关系
	if err := tx.Exec("DELETE FROM liked_commodity WHERE user_basic_id = ? AND commodity_basic_id = ?", user.ID, commodityID).Error; err != nil {
		// 处理错误
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "取消点赞失败"})
		return
	}

	// 更新商品的点赞数
	targetProduct.LikeCount--
	if err := tx.Save(&targetProduct).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "更新商品失败"})
		return
	}

	// 在本地切片中删除取消点赞的商品
	user.LikedCommodities = append(user.LikedCommodities[:likedIndex], user.LikedCommodities[likedIndex+1:]...)

	// 更新用户信息
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "更新用户信息失败"})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "事务提交失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "取消点赞成功"})
}

// UserCollectPro
// @Summary 收藏商品
// @Description 用户收藏商品
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param commodity_identity query string true "商品标识"
// @Success 200 {string} json {"code": 200, "msg": "收藏成功"}
// @Failure 401 {string} json {"code": 401, "message": "未授权"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 409 {string} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/collect/product [post]
func (u BasicOperateUser) UserCollectPro(c *gin.Context) {
	identity := c.Query("commodity_identity")
	id, err := strconv.Atoi(identity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "未授权"})
		return
	}

	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	_, existsuser, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	var pro models.CommodityBasic
	if err := dao.DB.Where("id = ?", uint(id)).First(&pro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	var product models.UserBasic
	if err := dao.DB.Preload("CollectedCommodities").Where("user_identity = ?", userClaims.UserIdentity).
		First(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// 检查用户是否已经收藏过该商品
	for _, collectedProduct := range product.CollectedCommodities {
		if collectedProduct.ID == uint(id) {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "已经收藏过该商品"})
			return
		}
	}

	tx := dao.DB.Begin()

	pro.CollectCount++
	pro.UpdateCommodity(pro)

	product.CollectedCommodities = append(product.CollectedCommodities, &pro)
	existsuser.CollectedCommodities = product.CollectedCommodities
	if err := existsuser.SaveUser(existsuser); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "提交事务失败"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "提交事务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "收藏成功"})
}

// UserGetCollectPro
// @Summary 获取用户收藏的商品数据
// @Description 获取用户收藏的商品数据
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} models.UserBasic "获取收藏的商品数据"
// @Failure 401 {string} json {"code": 401, "message": "未授权"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/get/collect/pro [get]
func (u BasicOperateUser) UserGetCollectPro(c *gin.Context) {
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
	var product models.UserBasic
	if err := dao.DB.Preload("CollectedCommodities").Where("user_identity = ?", userClaims.UserIdentity).
		First(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	var collectedCommodities []*models.CommodityBasic
	collectedCommodities = product.CollectedCommodities
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取收藏的商品数据",
		"data": collectedCommodities,
	})
} // UsersUncollectPro
// @Summary 取消收藏商品
// @Description 用户取消收藏商品
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param commodity_identity query string true "商品标识"
// @Success 200 {string} json {"code": 200, "msg": "取消收藏成功"}
// @Failure 401 {string} json {"code": 401, "message": "未授权"}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 409 {string} json {"code": 409, "msg": "该账号没有被注册"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/uncollect/product [post]
func (u BasicOperateUser) UsersUncollectPro(c *gin.Context) {
	identity := c.Query("commodity_identity")
	commodityID, err := strconv.Atoi(identity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "商品标识解析错误"})
		return
	}

	// 解析用户信息
	userClaim, exist := c.Get(pkg.UserClaimsContextKey)
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "未授权"})
		return
	}

	// 将 userClaim 转换为你的 UserClaims 结构体
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	// 检查用户是否已注册
	userExists, _, err := models.UserBasic{}.FindUserByAccountAndPassword(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	if !userExists {
		c.JSON(http.StatusConflict, gin.H{"code": http.StatusConflict, "msg": "该账号没有被注册"})
		return
	}

	// 开启数据库事务
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询用户信息及收藏的商品
	var user models.UserBasic
	if err := tx.Preload("CollectedCommodities").Where("user_identity = ?", userClaims.UserIdentity).
		First(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误"})
		return
	}

	if user.ID == 0 {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "未找到用户信息"})
		return
	}

	// 查询被取消收藏的商品
	var targetProduct models.CommodityBasic
	if err := tx.Where("id = ?", uint(commodityID)).First(&targetProduct).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "查询商品失败"})
		return
	}

	// 检查用户是否已经收藏过该商品
	var collected bool
	var collectedIndex int
	for i, collectedProduct := range user.CollectedCommodities {
		if collectedProduct.ID == uint(commodityID) {
			collected = true
			collectedIndex = i
			break
		}
	}

	if !collected {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "未收藏过该商品"})
		return
	}

	// 从关联表中删除收藏关系
	if err := tx.Exec("DELETE FROM collected_commodity WHERE user_basic_id = ? AND commodity_basic_id = ?", user.ID, commodityID).Error; err != nil {
		// 处理错误
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "取消收藏失败"})
		return
	}

	// 更新商品的收藏数
	targetProduct.CollectCount--
	if err := tx.Save(&targetProduct).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "更新商品失败"})
		return
	}

	// 在本地切片中删除取消收藏的商品
	user.CollectedCommodities = append(user.CollectedCommodities[:collectedIndex], user.CollectedCommodities[collectedIndex+1:]...)

	// 更新用户信息
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "更新用户信息失败"})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "事务提交失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "取消收藏成功"})
}
