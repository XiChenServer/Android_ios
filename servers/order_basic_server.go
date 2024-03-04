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
		log.Println("Error parsing productID:", err)
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
	//// 查询购买者的所有订单信息
	//var buyerOrders []models.Order
	//if err := dao.DB.Where("buyer_identity = ?", userClaims.UserIdentity).Find(&buyerOrders).Error; err != nil {
	//	log.Println("Error querying buyer's orders:", err)
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"code": 500,
	//		"msg":  "服务器内部错误",
	//	})
	//	return
	//}

	//// 遍历订单，检查是否存在相同的产品ID
	//for _, existingOrder := range buyerOrders {
	//	if existingOrder.ProductIdentity == uint(productID) {
	//		log.Println("Order already exists for the product.")
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"code": 400,
	//			"msg":  "订单已存在",
	//		})
	//		return
	//	}
	//}
	var product1 models.CommodityBasic
	var lock sync.Mutex
	lock.Lock()
	if err := dao.DB.Where("id = ?", productID).Find(&product1).Error; err != nil {
		log.Println("Error querying product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  "服务器内部错误",
		})
		return
	}
	if product1.Number < productNum {
		log.Println("Product does not belong to the seller.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "商品数量不够",
		})
		return
	}
	if product1.CommodityIdentity != seller.UserIdentity {
		log.Println("Product does not belong to the seller.")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	if product1.Number-productNum <= 0 {
		product1.SoldStatus = 3 // 已经卖完
	}
	product1.Number = product1.Number - productNum
	err = dao.DB.Updates(&product1).Error
	if err != nil {
		log.Println("Error updating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	lock.Unlock()
	if product1.SoldStatus == 1 {
		var order = &models.Order{
			OrderIdentity:   pkg.GenerateUniqueID(),
			BuyerIdentity:   userClaims.UserIdentity,
			SellerIdentity:  sellerIdentity,
			ProductIdentity: uint(productID),
			Name:            product1.Title,
			Price:           product1.Price,
			Quantity:        product1.Number,
			Msg:             product1.Information,
			Buyer:           buyer,
			Seller:          seller,
			Product:         product1,
		}

		err = models.Order{}.CreateOrder(order)
		if err != nil {
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

		err = dao.ProduceMessage(topic, message)
		if err != nil {
			log.Println("Error producing Kafka message: ", err)
			// 处理消息发送失败的情况
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "订单创建完成",
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

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

// FindSellOrderProduct 根据商品名称在用户订单中进行模糊匹配查询。
// @Summary 根据商品名称在用户订单中进行模糊匹配查询
// @Description 根据商品名称在用户订单中进行模糊匹配查询，并返回查询结果
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param name formData string true "商品名称"
// @Success 200 {string} json {"code":200,"msg":"找到以下数据","data":[]}
// @Failure 400 {string} json {"code":400,"msg":"没找到"}
// @Router /user/order/find/by_name [post]
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
	err = dao.DB.Where("seller_identity = ? AND (name LIKE ? or msg LIKE ?)", userClaims.UserIdentity, "%"+name+"%", "%"+name+"%").Find(&products).Error
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

// FindbuyOrderProduct 根据商品名称在用户订单中进行模糊匹配查询。
// @Summary 根据商品名称在用户订单中进行模糊匹配查询
// @Description 根据商品名称在用户订单中进行模糊匹配查询，并返回查询结果
// @Tags 订单
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param name formData string true "商品名称"
// @Success 200 {string} json {"code":200,"msg":"找到以下数据","data":[]}
// @Failure 400 {string} json {"code":400,"msg":"没找到"}
// @Router /user/order/find/by_name [post]
func (OrderBasicServer) FindbuyOrderProduct(c *gin.Context) {
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
	err = dao.DB.Where("seller_identity = ? AND (name LIKE ? or msg LIKE ?)", userClaims.UserIdentity, "%"+name+"%", "%"+name+"%").Find(&products).Error
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
