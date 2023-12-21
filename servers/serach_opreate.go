package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"github.com/gin-gonic/gin"
)

type SearchOperate struct {
}

func (SearchOperate) SearchProduct(c *gin.Context) {
	data := c.Query("data")
	var product []models.CommodityBasic
	if err := dao.DB.Where("title ")
}
