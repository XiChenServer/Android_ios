package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchOperate struct {
}

// SearchProduct
// @Summary 商品搜索
// @Description 根据关键字搜索商品信息
// @Tags 商品
// @Produce json
// @Param data query string true "搜索关键字"
// @Success 200 {string} json {"code": "200", "msg": "查找信息成功", "data": []models.CommodityBasic}
// @Failure 400 {string} json {"code": "400", "msg": "请求错误"}
// @Failure 401 {string} json {"code": "401", "msg": "未授权"}
// @Failure 404 {string} json {"code": "404", "msg": "商品不存在"}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /search/product [post]
func (SearchOperate) SearchProduct(c *gin.Context) {
	data := c.Query("data")
	var product []models.CommodityBasic
	if err := dao.DB.Where("title LIKE ? OR information LIKE ?", "%"+data+"%", "%"+data+"%").Find(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	if len(product) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "查找信息成功",
		"data": product,
	})
}
