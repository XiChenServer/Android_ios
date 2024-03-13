package router

import (
	"Android_ios/middleware"
	"Android_ios/servers"
	"fmt"
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

	r.StaticFS("/picture", gin.Dir("picture", true))
	// 使用 CORS 中间件处理跨域请求
	r.Use(cors.Default())

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.POST("/send_phone_code", servers.BasicServer{}.SendPhoneCode)
	r.POST("/send_email_code", servers.BasicServer{}.SendEmailCode)

	r.POST("/chat/basic/create", servers.UserChatBasicServer{}.CreateUserChat)

	r.POST("/user/register/phone", servers.BasicOperateUser{}.UserRegisterByPhone)

	r.POST("/user/login/phone", servers.BasicOperateUser{}.UserLoginByPhoneCode)
	r.POST("/user/login/password", servers.BasicOperateUser{}.UserLoginByPassword)
	r.POST("/user/login/phone_and_password", servers.BasicOperateUser{}.UserLoginByPhoneAndPassword)
	r.POST("/user/register/email", servers.BasicOperateUser{}.UserRegisterByEmail)
	r.POST("/user/login/email", servers.BasicOperateUser{}.UserLoginByEmailCode)
	r.POST("/user/login/email_and_password", servers.BasicOperateUser{}.UserLoginByEmailAndPassword)

	r.GET("/products/simple_info", servers.CommodityServer{}.GetProductsSimpleInfo)
	r.POST("/get/one_product_info", servers.CommodityServer{}.GetOneProAllInfo)
	r.POST("/search/product", servers.SearchOperate{}.SearchProduct)
	r.POST("/get/user_all_pro_list", servers.CommodityServer{}.GetUserAllProList)
	r.POST("/get/product/by_category", servers.CategoryServer{}.FindProByCategory)
	r.POST("/get/avatar/local", servers.BasicOperateUser{}.UserGetAvatarLocal)
	r.GET("/123", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "123",
		})
	})
	r.GET("/ws", func(c *gin.Context) {
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
		//servers.MsgHandler(c, ws)
	})
	r.GET("/auction/data", servers.BidRecordServer{}.FindAuctionData)
	user := r.Group("/user", middleware.AuthMiddleware())
	{

		user.GET("/find/all_prolist", servers.CommodityServer{}.UserFindAllProList)
		user.GET("/auction/info", servers.BidRecordServer{}.FindUserAuctionInfo)
		user.POST("/bid", servers.BidRecordServer{}.BidRecord)

		user.POST("/modify/password/by_email", servers.BasicOperateUser{}.ModifyUserPasswordByEmail)
		user.POST("/modify/password", servers.BasicOperateUser{}.ModifyUserPassword)
		user.POST("/order/find/buy_name", servers.OrderBasicServer{}.FindBuyOrderProduct)
		user.POST("/shopping_car/find", servers.CommodityServer{}.FindShoppingCarProduct)
		user.POST("/sell/find", servers.CommodityServer{}.FindSellProduct)
		user.POST("/order/find/sell_name", servers.OrderBasicServer{}.FindSellOrderProduct)

		user.POST("/uploads/background", servers.BasicOperateUser{}.UserUploadsBackground)
		user.POST("/shopping_car/delete", servers.CommodityServer{}.DelShoppingCar)
		user.GET("/shopping_car/view", servers.CommodityServer{}.ViewShoppingCar)
		user.POST("/shopping_car/add", servers.CommodityServer{}.AddShoppingCar)
		user.GET("/product/update_reco_prod_history", servers.CommodityServer{}.UpdateRecommendation)
		user.GET("/product/reco_prod_by_l_and_c", servers.CommodityServer{}.RecoProdByLAndC)

		user.POST("/modify/info", servers.BasicOperateUser{}.UserModifyInfo)
		user.POST("/upload/address", servers.BasicOperateUser{}.UserUploadAddress)
		user.POST("/uploads/avatar", servers.BasicOperateUser{}.UserUploadsAvatar)
		user.DELETE("/delete/address", servers.BasicOperateUser{}.UserDeleteAddress)
		user.POST("/deletes/products", servers.CommodityServer{}.UserDeletesProduct)
		user.POST("/changes/mobile/phone", servers.BasicOperateUser{}.UserChangesMobilePhoneNumber)
		user.POST("/adds/products", servers.CommodityServer{}.UserAddsProducts)
		user.GET("/get/avatar", servers.BasicOperateUser{}.UserGetAvatar)
		user.GET("/get/info", servers.BasicOperateUser{}.UserGetInfo)
		user.POST("/modifies/products", servers.CommodityServer{}.UserModifiesProducts)
		user.POST("/like/product", servers.BasicOperateUser{}.UsersLikePro)
		user.GET("/get/like/pro", servers.BasicOperateUser{}.UserGetLikePro)
		user.POST("/unlike/product", servers.BasicOperateUser{}.UsersUnlikePro)
		user.POST("/collect/product", servers.BasicOperateUser{}.UserCollectPro)
		user.GET("/get/collect/pro", servers.BasicOperateUser{}.UserGetCollectPro)
		user.POST("/uncollect/product", servers.BasicOperateUser{}.UsersUncollectPro)
		chat := user.Group("/chat")
		{
			chat.GET("/sendMsg", servers.UserChatServer{}.SendMsg)
			chat.GET("/sendUserMsg", servers.UserChatServer{}.SendUserMsg)
			chat.POST("/redisMsg", servers.UserChatServer{}.RedisMsg)
			r.POST("/upload", servers.Upload)
		}
		order := user.Group("/order")
		{
			order.POST("/find/buy/by_name", servers.OrderBasicServer{}.FindBuyOrderProduct)
			order.POST("/find/sell/by_name", servers.OrderBasicServer{}.FindSellOrderProduct)
			order.POST("/create", servers.OrderBasicServer{}.UserCreateOrder)
			order.POST("/delete", servers.OrderBasicServer{}.UserDeleteOrder)
			order.GET("/find/AllBuyOrder", servers.OrderBasicServer{}.FindAllBuyOrder)
			order.GET("/find/AllSellOrders", servers.OrderBasicServer{}.FindAllSellOrders)
			order.POST("/detail", servers.OrderBasicServer{}.FindOrderDetail)
		}
	}

	admin := r.Group("/admin")
	{
		admin.POST("/add_new_category_info", servers.AdminServer{}.AddNewCategoryInfo)
		admin.POST("/add_new_son_category_info", servers.AdminServer{}.AddNewSonCategoryInfo)
		admin.GET("/get/all_list_categories", servers.AdminServer{}.GetAllCategoryList)
	}

	return r
}
