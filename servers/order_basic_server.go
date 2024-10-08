package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type OrderBasicServer struct {
}

var (
	mu sync.Mutex
	//product models.CommodityBasic
	err error
)

// UserCreateOrder 创建订单接口
// @Summary 创建订单
// @Description 创建订单接口
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param sellerIdentity query string true "卖家身份"
// @Param productID query string true "商品ID"
// @Param productNum query string true "需要购买的商品数量"
// @Success 200 {string} json {"code":200,"msg":"订单创建完成"}
// @Failure 400 {string} json {"code":400,"msg":"卖家未拥有该商品"}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /user/order/create [post]
func (OrderBasicServer) UserCreateOrder(c *gin.Context) {
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	var buyer models.UserBasic
	if err := dao.DB.Where("user_identity = ?", userClaims.UserIdentity).Find(&buyer).Error; err != nil {
		log.Println("Error querying buyer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	sellerIdentity := c.Query("sellerIdentity")
	var seller models.UserBasic
	if err := dao.DB.Where("user_identity = ?", sellerIdentity).Find(&seller).Error; err != nil {
		log.Println("Error querying seller:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	productNum, err := strconv.Atoi(c.Query("productNum"))
	if err != nil {
		log.Println("Error parsing productNum:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的购买数量",
		})
		return
	}
	productID, err := strconv.Atoi(c.Query("productID"))
	if err != nil {
		log.Println("Error parsing productID:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的商品ID",
		})
		return
	}

	// 开启事务
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Recovered from panic:", r)
		}
	}()

	// 查询商品信息
	var product1 models.CommodityBasic
	if err := tx.Where("id = ?", productID).Find(&product1).Error; err != nil {
		tx.Rollback()
		log.Println("Error querying product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// 检查商品数量是否足够
	if product1.Number-productNum < 0 {
		tx.Rollback()
		log.Println("Insufficient quantity of product:", product1.Title)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "商品数量不足",
		})
		return
	}

	// 更新商品数量
	product1.Number -= productNum
	if product1.Number == 0 {
		product1.SoldStatus = 3
	}

	if err := tx.Save(&product1).Error; err != nil {
		tx.Rollback()
		log.Println("Error updating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "更新商品数量失败",
			"err":  err.Error(), // 返回具体的错误信息给客户端
		})
		return
	}

	// 创建订单
	var order = &models.Order{
		OrderIdentity:   pkg.GenerateUniqueID(),
		BuyerIdentity:   userClaims.UserIdentity,
		SellerIdentity:  sellerIdentity,
		ProductIdentity: uint(productID),
		Name:            product1.Title,
		Price:           product1.Price,
		Quantity:        productNum,
		Msg:             product1.Information,
		Buyer:           buyer,
		Seller:          seller,
		Product:         product1,
		Status:          1,
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		log.Println("Error creating order:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// 订单创建成功后发送消息到 Kafka
	topic := "order-created"
	message := fmt.Sprintf("Order created: %s", order.OrderIdentity)

	if err := dao.ProduceMessage(topic, message); err != nil {
		log.Println("Error producing Kafka message: ", err)
		// 处理消息发送失败的情况
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Println("Error committing transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "事务提交失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "订单创建完成",
	})
}

// UserDeleteOrder 删除订单接口
// @Summary 删除订单
// @Description 删除订单接口
// @Tags 订单
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param orderIdentity query string true "订单身份"
// @Success 200 {string} json {"code":200,"msg":"订单删除成功"}
// @Failure 400 {string} json {"code":400,"msg":"订单ID参数错误"}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /user/order/delete [post]
func (OrderBasicServer) UserDeleteOrder(c *gin.Context) {
	_, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	orderIdentity := c.Query("orderIdentity")
	// 查询订单
	var order models.Order
	if err := dao.DB.Where(" order_identity = ?", orderIdentity).Find(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	var product models.CommodityBasic
	if err := dao.DB.Where(" id = ?", order.ProductIdentity).Find(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	// 删除订单
	if err := dao.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	product.SoldStatus = 1
	err := dao.DB.Updates(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "订单删除成功",
	})
}

// FindAllBuyOrder 查询用户的所有购买订单
// @Summary 查询用户的所有购买订单
// @Description 查询用户的所有购买订单
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} json {"code":200,"msg":"成功获取购买订单","orders":{}}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /user/order/find/AllBuyOrder [get]
func (OrderBasicServer) FindAllBuyOrder(c *gin.Context) {
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	// 调用GetUserBuyOrders函数获取用户的所有购买订单
	orders, err := GetUserBuyOrders(userClaims.UserIdentity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "成功获取购买订单",
		"orders": orders,
	})
}

// GetUserBuyOrders 获取用户的所有购买订单
func GetUserBuyOrders(userID string) ([]models.Order, error) {
	var orders []models.Order

	err := dao.DB.Preload("Product").Preload("Buyer").Preload("Seller").
		Where("buyer_identity = ?", userID).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// FindAllSellOrders 查询所有出售订单
// @Summary 查询所有出售订单
// @Description 查询所有出售订单
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} json {"code":200,"msg":"成功获取所有出售订单","orders":{}}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /user/order/find/AllSellOrders [get]
func (OrderBasicServer) FindAllSellOrders(c *gin.Context) {
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	// 调用GetAllSellOrders函数获取所有出售订单
	orders, err := GetAllSellOrders(userClaims.UserIdentity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "成功获取所有出售订单",
		"orders": orders,
	})
}

// GetAllSellOrders 获取所有出售订单
func GetAllSellOrders(userID string) ([]models.Order, error) {
	var orders []models.Order
	err := dao.DB.Preload("Product").Preload("Buyer").Preload("Seller").
		Where("seller_identity = ?", userID).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// FindSellOrderProduct （出售订单）根据商品名称在用户订单中进行模糊匹配查询。
// @Summary （出售订单）根据商品名称在用户订单中进行模糊匹配查询
// @Description 根据商品名称在用户订单中进行模糊匹配查询，并返回查询结果
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param name formData string true "商品名称"
// @Success 200 {string} json {"code":200,"msg":"找到以下数据","data":[]}
// @Failure 400 {string} json {"code":400,"msg":"没找到"}
// @Router /user/order/find/sell/by_name [post]
func (OrderBasicServer) FindSellOrderProduct(c *gin.Context) {
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	name := c.PostForm("name")

	var products []models.Order
	err = dao.DB.Where("seller_identity = ? AND (name LIKE ? or msg LIKE ?)", userClaims.UserIdentity, "%"+name+"%", "%"+name+"%").Preload("Buyer").Preload("Seller").Preload("Product").Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "没找到",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "找到以下数据",
		"data": products,
	})
}

// FindBuyOrderProduct （购买订单）根据商品名称在用户订单中进行模糊匹配查询。
// @Summary （购买订单）根据商品名称在用户订单中进行模糊匹配查询
// @Description 根据商品名称在用户订单中进行模糊匹配查询，并返回查询结果
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param name formData string true "商品名称"
// @Success 200 {string} json {"code":200,"msg":"找到以下数据","data":[]}
// @Failure 400 {string} json {"code":400,"msg":"没找到"}
// @Router /user/order/find/buy/by_name [post]
func (OrderBasicServer) FindBuyOrderProduct(c *gin.Context) {
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	name := c.PostForm("name")

	var products []models.Order
	//err = dao.DB.Where("order_identity = ?", identity).Preload("Buyer").Preload("Seller").Preload("Product").Find(&products).Error

	err = dao.DB.Where("seller_identity = ? AND (name LIKE ? or msg LIKE ?)", userClaims.UserIdentity, "%"+name+"%", "%"+name+"%").Preload("Buyer").Preload("Seller").Preload("Product").Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "没找到",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "找到以下数据",
		"data": products,
	})
}

// FindOrderDetail 根据唯一标识查找详细信息。
// @Summary 根据唯一标识查找详细信息。
// @Description 根据商品名称在用户订单中进行模糊匹配查询，并返回查询结果
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param identity formData string identity "订单的唯一标识"
// @Success 200 {string} json {"code":200,"msg":"找到以下数据","data":[]}
// @Failure 400 {string} json {"code":400,"msg":"没找到"}
// @Router /user/order/detail [post]
func (OrderBasicServer) FindOrderDetail(c *gin.Context) {
	// 检查用户授权信息
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "未授权",
		})
		return
	}
	_, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	// 获取订单唯一标识
	identity := c.PostForm("identity")

	// 在数据库中查询订单信息
	var products models.Order
	//var products []models.Order
	err = dao.DB.Where("order_identity = ?", identity).Preload("Buyer").Preload("Seller").Preload("Product").Find(&products).Error
	//if err != nil {
	//	// 处理查询错误
	//}
	//
	//// products 变量现在包含了所有满足条件的订单信息
	//
	//err = dao.DB.Where("order_identity = ?", identity).Find(&products).Error
	if err != nil {
		// 如果未找到订单，返回 400 错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "没找到",
		})
		return
	}

	// 返回查询到的订单信息
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "找到以下数据",
		"data": products,
	})
}
