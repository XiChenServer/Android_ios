package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CategoryServer struct {
}

// FindProByCategory godoc
// @Summary 获取特定类别的所有商品
// @Description 根据kind_identity获取特定类别下的所有商品
// @Tags 商品
// @Produce json
// @Param kind_identity query string true "Kind Identity"
// @Success 200 {string} json {"code": 200, "msg": "获取该类型下的所有商品", "count": 10, "list": []}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /get/product/by_category [post]
func (CategoryServer) FindProByCategory(c *gin.Context) {
	identity := c.Query("kind_identity")
	var categories models.KindBasic
	var count int64
	if err := dao.DB.Unscoped().Preload("Commodities", func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at DESC") // 假设CommodityBasic有一个UpdatedAt字段
	}).Unscoped().
		Where("kind_basic.kind_identity = ?", identity).Order("updated_at desc").Find(&categories).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取该类型下的所有商品",
		"data": map[string]interface{}{
			"count": count,
			"list":  categories,
		},
	})
	return
}

// UserModifiesProducts
// @Summary 用户修改商品信息
// @Description 允许用户修改商品详细信息和媒体文件
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param commodity_id path string true "商品ID"
// @Param type formData []string true "商品类型"
// @Param title formData string true "商品标题"
// @Param number formData string true "商品数量"
// @Param information formData string true "商品信息"
// @Param price formData string true "商品价格"
// @Param is_auction formData string false "是否拍卖"
// @Param street formData string false "街道"
// @Param city formData string false "城市"
// @Param country formData string false "国家"
// @Param province formData string false "省份"
// @Param contact formData string false "联系人及其信息"
// @Param post_code formData string false "邮件地址"
// @Param files formData file true "商品图片"
// @Success 200 {string} json {"code": "200", "msg": "成功修改商品", "data": CommodityBasic}
// @Failure 400 {string} json {"code": "400", "msg": "请求无效。服务器无法理解请求"}
// @Failure 401 {string} json {"code": "401", "message": "未授权"}
// @Failure 404 {string} json {"code": "404", "msg": "商品不存在"}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /user/modifies/products/{commodity_id} [put]
