package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminServer struct {
}

// AddNewCategoryInfo godoc
// @Summary 添加新的分类
// @Description 根据传入的名称创建新的分类
// @Tags 管理员私有方法
// @Accept json
// @Produce json
// @Param name query string true "分类名称"
// @Success 200 {string} json {"code": 200, "msg": "新的分类创建完成"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /admin/add_new_category_info [post]
func (AdminServer) AddNewCategoryInfo(c *gin.Context) {
	var newCategory models.KindBasic
	name := c.Query("name")
	newCategory.Name = name
	newCategory.KindIdentity = pkg.GenerateUniqueID()
	newCategory.ParentID = 1
	// 将新的分类插入到数据库中
	result := dao.DB.Create(&newCategory)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器内部错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "新的分类创建完成"})
}

// AddNewSonCategoryInfo godoc
// @Summary 添加新的分类
// @Description 根据传入的名称创建新的分类
// @Tags 管理员私有方法
// @Accept json
// @Produce json
// @Param name query string true "分类名称"
// @Param identity query string true "父分类的identity"
// @Success 200 {string} json {"code": 200, "msg": "新的分类创建完成"}
// @Failure 404 {string} json {"code": 404, "msg": "父分类没有找到"}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /admin/add_new_son_category_info [post]
func (AdminServer) AddNewSonCategoryInfo(c *gin.Context) {
	var newCategory models.KindBasic
	identity := c.Query("identity")
	var parentCategory models.KindBasic
	if err := dao.DB.Where("kind_identity = ?", identity).Find(&parentCategory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "父分类没有找到"})
		return
	}

	name := c.Query("name")
	newCategory.Name = name

	// 设置 ParentID 为父分类的 ID
	newCategory.ParentID = parentCategory.ID

	// 将新的子分类的 ParentID 为父分类的 ID
	result := dao.DB.Create(&newCategory)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器内部错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "新的分类创建完成"})
}

// GetAllCategoryList godoc
// @Summary 查看所有分类
// @Description 获取系统中所有的分类信息
// @Tags 管理员私有方法
// @Accept json
// @Produce json
// @Success 200 {string} json {"code": 200, "msg": "获取分类信息成功", "data": [分类信息列表]}
// @Failure 500 {string} json {"code": 500, "msg": "服务器内部错误"}
// @Router /admin/get/all_list_categories [get]
func (AdminServer) GetAllCategoryList(c *gin.Context) {
	var categories []models.KindBasic
	result := dao.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器内部错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取分类信息成功", "data": categories})

}
