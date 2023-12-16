package router

import (
	"Android_ios/middleware"
	"Android_ios/servers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	// 使用 CORS 中间件
	r.Use(cors.Default())
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.POST("/send_phone_code", servers.BasicServer{}.SendPhoneCode)
	r.POST("/send_email_code", servers.BasicServer{}.SendEmailCode)

	r.POST("/user/register/phone", servers.BasicOperateUser{}.UserRegisterByPhone)

	r.POST("/user/login/phone", servers.BasicOperateUser{}.UserLoginByPhoneCode)
	r.POST("/user/login/password", servers.BasicOperateUser{}.UserLoginByPassword)
	r.POST("/user/login/phone_and_password", servers.BasicOperateUser{}.UserLoginByPhoneAndPassword)

	r.GET("/products/simple_info", servers.CommodityServer{}.GetProductsSimpleInfo)
	r.POST("/get/one_product_info", servers.CommodityServer{}.GetOneProAllInfo)
	r.POST("/get/user_all_pro_list", servers.CommodityServer{}.GetUserAllProList)
	r.POST("/get/product/by_category", servers.CategoryServer{}.FindProByCategory)

	user := r.Group("/user", middleware.AuthMiddleware())
	{
		user.PUT("/modify/info", servers.BasicOperateUser{}.UserModifyInfo)
		user.POST("/upload/address", servers.BasicOperateUser{}.UserUploadAddress)
		user.POST("/uploads/avatar", servers.BasicOperateUser{}.UserUploadsAvatar)
		user.DELETE("/delete/address", servers.BasicOperateUser{}.UserDeleteAddress)
		user.POST("/changes/mobile/phone", servers.BasicOperateUser{}.UserChangesMobilePhoneNumber)
		user.POST("/adds/products", servers.CommodityServer{}.UserAddsProducts)
		user.GET("/get/avatar", servers.BasicOperateUser{}.UserGetAvatar)
		user.GET("/get/info", servers.BasicOperateUser{}.UserGetInfo)
		user.POST("/modifies/products", servers.CommodityServer{}.UserModifiesProducts)
	}
	admin := r.Group("/admin")
	{
		admin.POST("/add_new_category_info", servers.AdminServer{}.AddNewCategoryInfo)
		admin.POST("/add_new_son_category_info", servers.AdminServer{}.AddNewSonCategoryInfo)
		admin.GET("/get/all_list_categories", servers.AdminServer{}.GetAllCategoryList)

	}
	return r
}

// CORSMiddleware 中间件处理CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
