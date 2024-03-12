package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type BidRecordServer struct {
}

// BidRecord godoc
// @Summary 用户进行竞拍
// @Description 用户进行竞拍操作，根据用户提交的竞拍信息更新竞拍记录或创建新的竞拍记录
// @Tags 拍卖
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param id formData string true "商品ID"
// @Param price formData string true "竞拍价格"
// @Success 200 {string} json {"code": 200, "msg": "bid placed successfully", "current_bid_info": {"commodity_identity": "商品ID", "user_id": "用户ID", "current_bid_price": "当前竞拍价格", "current_bid_time": "当前竞拍时间"}}
// @Failure 400 {string} json {"code": 400, "msg": "无效的价格"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 404 {string} json {"code": 404, "msg": "商品未找到或不在竞拍中"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /user/bid [post]
func (BidRecordServer) BidRecord(c *gin.Context) {
	// 解析用户信息
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "未授权",
		})
		return
	}

	// 将 userClaim 转换为您的 UserClaims 结构体
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}

	userIdStr := userClaims.UserIdentity
	productId, _ := strconv.Atoi(c.PostForm("id"))
	priceStr := c.PostForm("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	// 查询拍卖商品信息
	var product models.CommodityBasic
	err = dao.DB.Where("id = ? AND sold_status = ?", productId, 2).First(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "product not found or not in auction"})
		return
	}

	// 计算拍卖时间是否已结束
	auctionEndTime := product.CreatedAt.Add(10 * time.Minute)
	now := time.Now()
	if now.After(auctionEndTime) {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "auction has ended"})
		return
	}

	// 查询有没有之前的竞拍记录
	var bidRecord models.BidRecord
	err = dao.DB.Where("product_id = ? AND user_id = ?", productId, userIdStr).First(&bidRecord).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "failed to query bid record"})
			return
		}
		// 如果没有找到之前的竞拍记录，直接创建新的竞拍记录
		bidRecord = models.BidRecord{
			ProductId:   product.ID,
			UserID:      userIdStr,
			UserAccount: userClaims.Account,
			Price:       price,
			BidTime:     now,
		}
	} else {
		// 如果找到了之前的竞拍记录，并且当前出价高于之前的出价，则更新竞拍记录
		if price >= bidRecord.Price {
			bidRecord.Price = price
			bidRecord.BidTime = now
		}
	}

	// 创建或更新竞拍记录
	if err := dao.DB.Save(&bidRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "failed to create or update bid record"})
		return
	}

	// 创建当前竞拍信息记录
	currentBidInfo := models.UserBidRecord{
		ProductId:       product.ID,
		UserID:          userIdStr,
		CurrentBidPrice: bidRecord.Price,
		CurrentBidTime:  bidRecord.BidTime,
	}
	if err = dao.DB.Create(&currentBidInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "竞拍失败"})
		return
	}
	// 返回更新后的商品信息和当前竞拍信息给前端
	c.JSON(http.StatusOK, gin.H{
		"code":             200,
		"msg":              "bid placed successfully",
		"current_bid_info": currentBidInfo,
	})
}

// FindAuctionData godoc
// @Summary 获取拍卖数据
// @Description 查询拍卖中的商品以及检查是否满足生成订单的条件，并根据情况生成订单并删除商品
// @Tags 拍卖
// @Produce json
// @Success 200 {string} json {"code": 200, "msg": "拍卖数据如下", "data": [{"commodity_identity": "商品ID", "title": "商品标题", "number": "商品数量", "information": "商品信息", "price": "商品价格", "sold_status": "售出状态", "created_at": "创建时间", "updated_at": "更新时间"}, ...]}
// @Failure 500 {string} json {"code": 500, "msg": "无法获取拍卖数据或查询竞拍信息失败"}
// @Router /auction/data [get]
func (BidRecordServer) FindAuctionData(c *gin.Context) {
	var products []models.CommodityBasic

	// 计算拍卖结束时间（即创建时间加上10分钟）
	endTime := time.Now().Add(-10 * time.Minute)

	// 查询拍卖中且拍卖时间尚未到达的商品，并连接 BidRecord 表
	err := dao.DB.
		Select("commodity_basic.*, bid_record.*").
		Where("commodity_basic.sold_status = ? AND commodity_basic.created_at > ?", 2, endTime).
		Joins("LEFT JOIN bid_record ON bid_record.product_id = commodity_basic.id").
		Find(&products).Error

	if err != nil {
		handleError(c, "无法获取拍卖数据", http.StatusInternalServerError)
		return
	}

	endProducts := getEndAuctionProducts(endTime)

	if len(endProducts) != 0 {
		createOrdersAndDeleteProducts(c, endProducts)
	}

	// 返回拍卖数据给前端
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "拍卖数据如下",
		"data": products,
	})
}

func getEndAuctionProducts(endTime time.Time) []models.CommodityBasic {
	var endProducts []models.CommodityBasic

	// 查询拍卖中且拍卖时间已到达的商品
	err := dao.DB.Where("sold_status = ? AND created_at < ?", 2, endTime).Find(&endProducts).Error
	if err != nil {
		// 可以将日志记录放在这里
		return nil
	}

	return endProducts
}

func createOrdersAndDeleteProducts(c *gin.Context, endProducts []models.CommodityBasic) {
	for _, endProduct := range endProducts {
		err := createOrderIfBidHigher(&endProduct)
		if err != nil {
			handleError(c, "无法创建订单或删除商品", http.StatusInternalServerError)
			return
		}
	}
}
func createOrderIfBidHigher(product *models.CommodityBasic) error {
	// 查询该商品的竞拍记录
	var bidInfo models.BidRecord
	err := dao.DB.Where("product_id = ?", product.ID).First(&bidInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有人竞拍，直接删除商品
			if err := dao.DB.Delete(&product).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// 检查是否满足生成订单的条件（例如价格高于起拍价等）
	if bidInfo.Price > product.Price {
		// 创建订单
		order := &models.Order{
			OrderIdentity:   pkg.GenerateUniqueID(),    // 生成唯一订单号
			BuyerIdentity:   bidInfo.UserID,            // 竞拍者的用户ID
			SellerIdentity:  product.CommodityIdentity, // 商品所有者的用户ID
			ProductIdentity: product.ID,                // 商品ID
			Name:            product.Title,             // 商品名称
			Price:           bidInfo.Price,             // 成交价格为竞拍者的出价
			Quantity:        product.Number,            // 数量为商品的数量
			Msg:             product.Information,       // 商品信息
			Status:          1,                         // 订单状态为待支付
		}
		// 保存订单
		if err = dao.DB.Create(order).Error; err != nil {
			return err
		}
		// 订单创建成功后发送消息到 Kafka
		if err := notifyOrderCreation(order); err != nil {
			// 处理消息发送失败的情况
			return err
		}

		// 创建订单成功后，删除商品
		if err := dao.DB.Delete(&product).Error; err != nil {
			return err
		}
	}

	return nil
}

func handleError(c *gin.Context, msg string, code int) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func notifyOrderCreation(order *models.Order) error {
	topic := "order-created"
	message := fmt.Sprintf("Order created: %s", order.OrderIdentity)
	return dao.ProduceMessage(topic, message)
}

// FindUserAuctionInfo godoc
// @Summary 查找用户的拍卖记录
// @Description 查询特定用户的拍卖记录
// @Tags 拍卖
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} json {"code": 200, "msg": "数据如下", "data": [{"commodity_identity": "商品ID", "user_id": "用户ID", "current_bid_price": "当前竞拍价格", "current_bid_time": "当前竞拍时间"}, ...]}
// @Failure 400 {string} json {"code": 400, "msg": "请求错误"}
// @Failure 401 {string} json {"code": 401, "msg": "未授权"}
// @Failure 404 {string} json {"code": 404, "msg": "没有找到数据"}
// @Router /user/auction/info [get]
// 查找这个用户的拍卖记录
func (BidRecordServer) FindUserAuctionInfo(c *gin.Context) {
	// 解析用户信息
	userClaim, exists := c.Get(pkg.UserClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "未授权",
		})
		return
	}

	// 将 userClaim 转换为您的 UserClaims 结构体
	userClaims, ok := userClaim.(*pkg.UserClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "请求错误"})
		return
	}
	var userBidRecord []models.UserBidRecord

	err = dao.DB.Where("user_id = ?", userClaims.UserIdentity).Find(&userBidRecord).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "没有找到数据",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "数据如下",
		"data": userBidRecord,
	})

}
