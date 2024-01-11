package router

import (
	"Android_ios/middleware"
	"Android_ios/servers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"time"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 将 /picture 路径下的所有文件映射到路由上
	r.StaticFS("/picture", gin.Dir("picture", true))
	//	r.Use(CORSMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:13000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.POST("/send_phone_code", servers.BasicServer{}.SendPhoneCode)
	r.POST("/send_email_code", servers.BasicServer{}.SendEmailCode)

	r.POST("/user/register/phone", servers.BasicOperateUser{}.UserRegisterByPhone)

	r.POST("/user/login/phone", servers.BasicOperateUser{}.UserLoginByPhoneCode)
	r.POST("/user/login/password", servers.BasicOperateUser{}.UserLoginByPassword)
	r.POST("/user/login/phone_and_password", servers.BasicOperateUser{}.UserLoginByPhoneAndPassword)

	r.GET("/products/simple_info", servers.CommodityServer{}.GetProductsSimpleInfo)
	r.POST("/get/one_product_info", servers.CommodityServer{}.GetOneProAllInfo)
	r.POST("/search/product", servers.SearchOperate{}.SearchProduct)
	r.POST("/get/user_all_pro_list", servers.CommodityServer{}.GetUserAllProList)
	r.POST("/get/product/by_category", servers.CategoryServer{}.FindProByCategory)
	r.POST("/get/avatar/local", servers.BasicOperateUser{}.UserGetAvatarLocal)
	//r.POST("/user/like/product ", servers.BasicOperateUser{}.UsersLikePro)
	user := r.Group("/user", middleware.AuthMiddleware())
	{
		user.PUT("/modify/info", servers.BasicOperateUser{}.UserModifyInfo)
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
		//user.POST("/upload/local", servers.BasicOperateUser{}.UserUploadLocal)

	}
	admin := r.Group("/admin")
	{
		admin.POST("/add_new_category_info", servers.AdminServer{}.AddNewCategoryInfo)
		admin.POST("/add_new_son_category_info", servers.AdminServer{}.AddNewSonCategoryInfo)
		admin.GET("/get/all_list_categories", servers.AdminServer{}.GetAllCategoryList)
	}
	return r
}

//// CORSMiddleware 中间件处理CORS
//func CORSMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
//
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//
//		c.Next()
//	}
//}
