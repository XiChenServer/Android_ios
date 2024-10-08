package servers

import (
	"Android_ios/models"
	"Android_ios/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserChatServer struct {
}

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
//func (UserChatServer) GetUserList(c *gin.Context) {
//	data := make([]*models.UserBasic, 0)
//	data = models.UserChatBasic{}.GetUserList()
//	c.JSONP(http.StatusOK, gin.H{
//		"message": data,
//	})
//}

//// CreateUser
//// @Summary 新增用户
//// @Tags 用户模块
//// @param name query string false "用户名"
//// @param password query string false "密码"
//// @param repassword query string false "确认密码"
//// @Success 200 {string} json{"code","message"}
//// @Router /user/createUser [get]
//func (UserChatServer) CreateUser(c *gin.Context) {
//
//	// user.Name = c.Query("name")
//	// password := c.Query("password")
//	// repassword := c.Query("repassword")
//	user := models.UserChatBasic{}
//	user.Name = c.Request.FormValue("name")
//	password := c.Request.FormValue("password")
//	repassword := c.Request.FormValue("Identity")
//	fmt.Println(user.Name, "  >>>>>>>>>>>  ", password, repassword)
//	salt := fmt.Sprintf("%06d", rand.Int31())
//
//	data := models.UserChatBasic{}.FindUserByName(user.Name)
//	if user.Name == "" || password == "" || repassword == "" {
//		c.JSON(200, gin.H{
//			"code":    -1, //  0成功   -1失败
//			"message": "用户名或密码不能为空！",
//			"data":    user,
//		})
//		return
//	}
//	if data.Account != "" {
//		c.JSON(200, gin.H{
//			"code":    -1, //  0成功   -1失败
//			"message": "用户名已注册！",
//			"data":    user,
//		})
//		return
//	}
//	if password != repassword {
//		c.JSON(200, gin.H{
//			"code":    -1, //  0成功   -1失败
//			"message": "两次密码不一致！",
//			"data":    user,
//		})
//		return
//	}
//	//user.PassWord = password
//	user.PassWord = utils.MakePassword(password, salt)
//	user. = salt
//	fmt.Println(user.PassWord)
//	user.LoginTime = time.Now()
//	user.LoginOutTime = time.Now()
//	user.HeartbeatTime = time.Now()
//	models.CreateUser(user)
//	c.JSON(200, gin.H{
//		"code":    0, //  0成功   -1失败
//		"message": "新增用户成功！",
//		"data":    user,
//	})
//}
//
//// FindUserByNameAndPwd
//// @Summary 所有用户
//// @Tags 用户模块
//// @param name query string false "用户名"
//// @param password query string false "密码"
//// @Success 200 {string} json{"code","message"}
//// @Router /user/findUserByNameAndPwd [post]
//func (UserChatServer) FindUserByNameAndPwd(c *gin.Context) {
//	data := models.UserBasic{}
//
//	//name := c.Query("name")
//	//password := c.Query("password")
//	name := c.Request.FormValue("name")
//	password := c.Request.FormValue("password")
//	fmt.Println(name, password)
//	user := models.UserChatBasic{}.FindUserByName(name)
//	if user.NickName == "" {
//		c.JSON(200, gin.H{
//			"code":    -1, //  0成功   -1失败
//			"message": "该用户不存在",
//			"data":    data,
//		})
//		return
//	}
//
//	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
//	if !flag {
//		c.JSON(200, gin.H{
//			"code":    -1, //  0成功   -1失败
//			"message": "密码不正确",
//			"data":    data,
//		})
//		return
//	}
//	pwd := utils.MakePassword(password, user.Salt)
//	data = models.FindUserByNameAndPwd(name, pwd)
//
//	c.JSON(200, gin.H{
//		"code":    0, //  0成功   -1失败
//		"message": "登录成功",
//		"data":    data,
//	})
//}

// // DeleteUser
// // @Summary 删除用户
// // @Tags 用户模块
// // @param id query string false "id"
// // @Success 200 {string} json{"code","message"}
// // @Router /user/deleteUser [get]
//
//	func (UserChatServer) DeleteUser(c *gin.Context) {
//		user := models.UserBasic{}
//		id, _ := strconv.Atoi(c.Query("id"))
//		user.ID = uint(id)
//		models.DeleteUser(user)
//		c.JSON(200, gin.H{
//			"code":    0, //  0成功   -1失败
//			"message": "删除用户成功！",
//			"data":    user,
//		})
//
// }
//
// // UpdateUser
// // @Summary 修改用户
// // @Tags 用户模块
// // @param id formData string false "id"
// // @param name formData string false "name"
// // @param password formData string false "password"
// // @param phone formData string false "phone"
// // @param email formData string false "email"
// // @Success 200 {string} json{"code","message"}
// // @Router /user/updateUser [post]
//
//	func (UserChatServer) UpdateUser(c *gin.Context) {
//		user := models.UserBasic{}
//		id, _ := strconv.Atoi(c.PostForm("id"))
//		user.ID = uint(id)
//		user.Name = c.PostForm("name")
//		user.PassWord = c.PostForm("password")
//		user.Phone = c.PostForm("phone")
//		user.Avatar = c.PostForm("icon")
//		user.Email = c.PostForm("email")
//		fmt.Println("update :", user)
//
//		_, err := govalidator.ValidateStruct(user)
//		if err != nil {
//			fmt.Println(err)
//			c.JSON(200, gin.H{
//				"code":    -1, //  0成功   -1失败
//				"message": "修改参数不匹配！",
//				"data":    user,
//			})
//		} else {
//			models.UserChatBasic{}.UpdateUser(user)
//			c.JSON(200, gin.H{
//				"code":    0, //  0成功   -1失败
//				"message": "修改用户成功！",
//				"data":    user,
//			})
//		}
//
// }
// 防止跨域站点伪造请求
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMsg
// @Summary 通过WebSocket发送消息
// @Description 建立WebSocket连接并发送消息。
// @ID sendWebSocketMessage
// @Tags 用户私有方法
// @Produce plain
// @Success 101 {string} string "WebSocket连接已建立"
// @Failure 403 {string} string "WebSocket连接失败"
// @Router /user/chat/sendMsg [get]
func (UserChatServer) SendMsg(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(c, ws)
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := pkg.Subscribe(c, pkg.PublishKey)
		if err != nil {
			fmt.Println(" MsgHandler 发送失败", err)
		}

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// RedisMsg
// @Summary 从Redis获取用户聊天消息
// @Description 从Redis中获取两个用户之间的聊天消息。
// @Tags 用户私有方法
// @ID getRedisMessages
// @Accept json
// @Produce json
// @Param userIdA formData int true "用户A的ID"
// @Param userIdB formData int true "用户B的ID"
// @Param start formData int true "消息起始位置"
// @Param end formData int true "消息结束位置"
// @Param isRev formData bool true "是否反转消息顺序"
// @Success 200 {string} ResponseObject "成功返回"
// @Router /user/chat/redisMsg [post]
func (UserChatServer) RedisMsg(c *gin.Context) {
	userIdA, _ := strconv.Atoi(c.PostForm("userIdA"))
	userIdB, _ := strconv.Atoi(c.PostForm("userIdB"))
	start, _ := strconv.Atoi(c.PostForm("start"))
	end, _ := strconv.Atoi(c.PostForm("end"))
	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))
	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	pkg.RespOKList(c.Writer, "ok", res)
}

func (UserChatServer) MsgHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := pkg.Subscribe(c, pkg.PublishKey)
		if err != nil {
			fmt.Println(" MsgHandler 发送失败", err)
		}

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// SendUserMsg
// @Summary 发送用户消息
// @Description 通过WebSocket向指定用户发送消息。
// @Tags 用户私有方法
// @ID sendUserMsg
// @Accept json
// @Produce json
// @Param userId query int true "接收者的用户ID"
// @Router /user/chat/sendUserMsg [get]
func (UserChatServer) SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func (UserChatServer) SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(uint(id))
	// c.JSON(200, gin.H{
	// 	"code":    0, //  0成功   -1失败
	// 	"message": "查询好友列表成功！",
	// 	"data":    users,
	// })
	pkg.RespOKList(c.Writer, users, len(users))
}

func (UserChatServer) AddFriend(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	targetName := c.Request.FormValue("targetName")
	//targetId, _ := strconv.Atoi(c.Request.FormValue("targetId"))
	code, msg := models.AddFriend(uint(userId), targetName)
	if code == 0 {
		pkg.RespOK(c.Writer, code, msg)
	} else {
		pkg.RespFail(c.Writer, msg)
	}
}

//// 新建群
//func (UserChatServer) CreateCommunity(c *gin.Context) {
//	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
//	name := c.Request.FormValue("name")
//	icon := c.Request.FormValue("icon")
//	desc := c.Request.FormValue("desc")
//	community := models.Community{}
//	community.OwnerId = uint(ownerId)
//	community.Name = name
//	community.Img = icon
//	community.Desc = desc
//	code, msg := models.CreateCommunity(community)
//	if code == 0 {
//		utils.RespOK(c.Writer, code, msg)
//	} else {
//		utils.RespFail(c.Writer, msg)
//	}
//}

//// 加载群列表
//func (UserChatServer) LoadCommunity(c *gin.Context) {
//	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
//	//	name := c.Request.FormValue("name")
//	data, msg := models.LoadCommunity(uint(ownerId))
//	if len(data) != 0 {
//		pkg.RespList(c.Writer, 0, data, msg)
//	} else {
//		pkg.RespFail(c.Writer, msg)
//	}
//}
//
//// 加入群 userId uint, comId uint
//func (UserChatServer) JoinGroups(c *gin.Context) {
//	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
//	comId := c.Request.FormValue("comId")
//
//	//	name := c.Request.FormValue("name")
//	data, msg := models.UserChatBasic{}.JoinGroup(uint(userId), comId)
//	if data == 0 {
//		pkg.RespOK(c.Writer, data, msg)
//	} else {
//		pkg.RespFail(c.Writer, msg)
//	}
//}

func (UserChatServer) FindByID(c *gin.Context) {
	//userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	account := c.Request.FormValue("account")
	//	name := c.Request.FormValue("name")
	data := models.UserChatBasic{}.FindByAccount(account)
	pkg.RespOK(c.Writer, data, "ok")
}
