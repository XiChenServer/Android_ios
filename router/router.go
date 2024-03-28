package router

import (
	"Android_ios/middleware"
	"Android_ios/servers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Router() *gin.Engine {
	r := gin.Default()

	// 静态文件服务，用于提供静态图片文件的访问服务
	r.StaticFS("/picture", gin.Dir("picture", true))

	// 使用CORS中间件处理跨域请求
	r.Use(cors.Default())

	// 访问Swagger API文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 手机验证码接口
	r.POST("/send_phone_code", servers.BasicServer{}.SendPhoneCode)

	// 邮箱验证码接口
	r.POST("/send_email_code", servers.BasicServer{}.SendEmailCode)

	// 聊天信息创建接口
	r.POST("/chat/basic/create", servers.UserChatBasicServer{}.CreateUserChat)

	// 手机号注册接口
	r.POST("/user/register/phone", servers.BasicOperateUser{}.UserRegisterByPhone)

	// 用户登录接口
	r.POST("/user/login/phone", servers.BasicOperateUser{}.UserLoginByPhoneCode)
	r.POST("/user/login/password", servers.BasicOperateUser{}.UserLoginByPassword)
	r.POST("/user/login/phone_and_password", servers.BasicOperateUser{}.UserLoginByPhoneAndPassword)
	r.POST("/user/register/email", servers.BasicOperateUser{}.UserRegisterByEmail)
	r.POST("/user/login/email", servers.BasicOperateUser{}.UserLoginByEmailCode)
	r.POST("/user/login/email_and_password", servers.BasicOperateUser{}.UserLoginByEmailAndPassword)

	// 获取商品信息接口
	// 获取商品简单信息
	r.GET("/products/simple_info", servers.CommodityServer{}.GetProductsSimpleInfo)
	// 获取一个商品的信息
	r.POST("/get/one_product_info", servers.CommodityServer{}.GetOneProAllInfo)
	// 搜索商品
	r.POST("/search/product", servers.SearchOperate{}.SearchProduct)
	// 获取一个用户的所有商品
	r.POST("/get/user_all_pro_list", servers.CommodityServer{}.GetUserAllProList)

	r.POST("/get/product/by_category", servers.CategoryServer{}.FindProByCategory)
	r.POST("/get/avatar/local", servers.BasicOperateUser{}.UserGetAvatarLocal)

	// 123测试接口
	r.GET("/123", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "123",
		})
	})

	// WebSocket接口
	r.GET("/ws", func(c *gin.Context) {
		// 处理WebSocket连接
	})

	// 拍卖数据接口
	r.GET("/auction/data", servers.BidRecordServer{}.FindAuctionData)

	// 用户相关操作路由组
	user := r.Group("/user", middleware.AuthMiddleware())
	{
		// 用户商品列表接口
		user.GET("/find/all_prolist", servers.CommodityServer{}.UserFindAllProList)

		// 用户竞拍信息接口
		user.GET("/auction/info", servers.BidRecordServer{}.FindUserAuctionInfo)

		// 用户出价接口
		user.POST("/bid", servers.BidRecordServer{}.BidRecord)

		// 用户修改密码接口
		user.POST("/modify/password/by_email", servers.BasicOperateUser{}.ModifyUserPasswordByEmail)
		user.POST("/modify/password", servers.BasicOperateUser{}.ModifyUserPassword)

		// 用户订单相关接口
		user.POST("/order/find/buy_name", servers.OrderBasicServer{}.FindBuyOrderProduct)
		// 购物车查找
		user.POST("/shopping_car/find", servers.CommodityServer{}.FindShoppingCarProduct)
		// 出售商品的查找
		user.POST("/sell/find", servers.CommodityServer{}.FindSellProduct)
		user.POST("/order/find/sell_name", servers.OrderBasicServer{}.FindSellOrderProduct)

		// 用户上传背景接口
		user.POST("/uploads/background", servers.BasicOperateUser{}.UserUploadsBackground)

		// 用户购物车相关接口
		user.POST("/shopping_car/delete", servers.CommodityServer{}.DelShoppingCar)
		// 查看购物车
		user.GET("/shopping_car/view", servers.CommodityServer{}.ViewShoppingCar)
		// 添加购物车
		user.POST("/shopping_car/add", servers.CommodityServer{}.AddShoppingCar)

		// 用户商品操作接口
		user.POST("/deletes/products", servers.CommodityServer{}.UserDeletesProduct)
		// 添加商品
		user.POST("/adds/products", servers.CommodityServer{}.UserAddsProducts)
		// 修改商品
		user.POST("/modifies/products", servers.CommodityServer{}.UserModifiesProducts)

		// 用户信息管理接口
		user.POST("/modify/info", servers.BasicOperateUser{}.UserModifyInfo)
		// 上传地址
		user.POST("/upload/address", servers.BasicOperateUser{}.UserUploadAddress)
		// 上传头像
		user.POST("/uploads/avatar", servers.BasicOperateUser{}.UserUploadsAvatar)
		// 删除地址
		user.DELETE("/delete/address", servers.BasicOperateUser{}.UserDeleteAddress)
		// 换绑手机号
		user.POST("/changes/mobile/phone", servers.BasicOperateUser{}.UserChangesMobilePhoneNumber)

		// 用户喜欢和收藏商品接口
		user.POST("/like/product", servers.BasicOperateUser{}.UsersLikePro)

		user.GET("/get/like/pro", servers.BasicOperateUser{}.UserGetLikePro)
		user.POST("/unlike/product", servers.BasicOperateUser{}.UsersUnlikePro)
		user.POST("/collect/product", servers.BasicOperateUser{}.UserCollectPro)
		user.GET("/get/collect/pro", servers.BasicOperateUser{}.UserGetCollectPro)
		user.POST("/uncollect/product", servers.BasicOperateUser{}.UsersUncollectPro)

		// 用户聊天相关接口
		chat := user.Group("/chat")
		{
			chat.GET("/sendMsg", servers.UserChatServer{}.SendMsg)
			chat.GET("/sendUserMsg", servers.UserChatServer{}.SendUserMsg)
			chat.POST("/redisMsg", servers.UserChatServer{}.RedisMsg)
			r.POST("/upload", servers.Upload)
		}

		// 用户订单相关接口
		order := user.Group("/order")
		{

			order.POST("/find/buy/by_name", servers.OrderBasicServer{}.FindBuyOrderProduct)
			order.POST("/find/sell/by_name", servers.OrderBasicServer{}.FindSellOrderProduct)
			// 创建订单
			order.POST("/create", servers.OrderBasicServer{}.UserCreateOrder)
			// 删除订单
			order.POST("/delete", servers.OrderBasicServer{}.UserDeleteOrder)
			order.GET("/find/AllBuyOrder", servers.OrderBasicServer{}.FindAllBuyOrder)
			order.GET("/find/AllSellOrders", servers.OrderBasicServer{}.FindAllSellOrders)
			// 订单细节
			order.POST("/detail", servers.OrderBasicServer{}.FindOrderDetail)
		}
	}

	// 管理员相关操作路由组
	admin := r.Group("/admin")
	{
		// 添加新分类信息接口
		admin.POST("/add_new_category_info", servers.AdminServer{}.AddNewCategoryInfo)
		admin.POST("/add_new_son_category_info", servers.AdminServer{}.AddNewSonCategoryInfo)
		admin.GET("/get/all_list_categories", servers.AdminServer{}.GetAllCategoryList)
	}

	return r
}
