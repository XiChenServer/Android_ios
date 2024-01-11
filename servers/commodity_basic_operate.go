package servers

import (
	"Android_ios/dao"
	"Android_ios/models"
	"Android_ios/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"net/http"
	"strconv"
)

type CommodityBasic struct {
	Type        []string `json:"type"`        //类型
	Title       string   `json:"title"`       // 商品的标题或名称，类型为varchar(36)
	Number      int      `json:"number"`      // 商品的数量或数量
	Information string   `json:"information"` // 有关商品的额外详细信息，类型为text
	Price       float64  `json:"price"`       // 商品的价格，类型为decimal(10,2)
	IsAuction   int      `json:"is_auction"`  // 商品是否属于拍卖
}

type CommodityServer struct {
}

// UserAddsProducts
// @Summary 用户添加商品
// @Description 允许用户添加具有详细信息和媒体文件的商品
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
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
// @Success 200 {string} json {"code": "200", "msg": "成功创建商品", "data": CommodityBasic}
// @Failure 400 {string} json {"code": "400", "msg": "请求无效。服务器无法理解请求"}
// @Failure 401 {string} json {"code": "401", "message": "未授权"}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /user/adds/products [post]
func (CommodityServer) UserAddsProducts(c *gin.Context) {
	// Parse user information
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
	// Get form-data values
	types := c.PostFormArray("type")
	title := c.PostForm("title")
	number := c.PostForm("number")
	information := c.PostForm("information")
	price := c.PostForm("price")
	isAuction := c.PostForm("is_auction")
	street := c.PostForm("street")
	city := c.PostForm("city")
	country := c.PostForm("country")
	province := c.PostForm("province")
	contact := c.PostForm("contact")
	postCode := c.PostForm("post_code")

	// Parse number and price
	// Parse number and price
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的参数：商品数量",
		})
		return
	}

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的参数：商品价格",
		})
		return
	}

	// Parse isAuction, handle empty value
	var isAuctionInt int
	if isAuction != "" {
		isAuctionInt, err = strconv.Atoi(isAuction)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "无效的参数：是否拍卖",
			})
			return
		}
	} else {
		// Set a default value or handle it as needed
		isAuctionInt = 0 // Assuming 0 as a default value, update as per your requirement
	}

	// Validate and associate types
	categories, err := validateAndAssociateTypes(types)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	// Create address entry
	modelAddr := models.AddressEntry{
		Street:   street,
		City:     city,
		Country:  country,
		Province: province,
		Contact:  contact,
		PostCode: postCode,
		Identity: pkg.GetAccountNumber(),
	}

	// Find or create user
	exists, existsuser, err := models.UserBasic{}.FindUserByAccount(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// Create product
	status := 0
	if isAuctionInt != status {
		status = 1
	}
	addrJson := models.JSONAddress{modelAddr}
	product := &models.CommodityBasic{
		CommodityIdentity: existsuser.UserIdentity,
		Title:             title,
		Number:            numberInt,
		Information:       information,
		Price:             priceFloat,
		SoldStatus:        status,
		Media:             nil,
		IsAuction:         isAuctionInt,
		Address:           addrJson,
		LikeCount:         0,
		CollectCount:      0,
		Categories:        categories,
	}

	// Upload files to OSS
	form, err := c.MultipartForm()
	// Get form-data values
	// ...
	files := form.File["files"] // 注意这里的字段名要与 curl 请求中的一致
	// ...

	for _, file := range files {
		// 保存文件到本地
		localFilePath := "./picture/commodity/" + file.Filename
		fmt.Println("sdf", localFilePath)
		err = c.SaveUploadedFile(file, localFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		objectKey := "picture/commodity" + file.Filename
		err = pkg.UploadAllFile(objectKey, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fileURL := fmt.Sprintf("http://%s", objectKey)
		fmt.Println(fileURL)
		media := models.MediaBasic{Image: fileURL}
		product.Media = append(product.Media, media)
	}

	// Save user and product
	existsuser.Commodity = append(existsuser.Commodity, product)
	if existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}
	if product.CreateCommodity(product); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "成功创建商品",
		"data": product,
	})
}

// validateAndAssociateTypes validates and associates product types.
// validateAndAssociateTypes 验证并关联商品类型。
func validateAndAssociateTypes(types []string) ([]*models.KindBasic, error) {
	var categories []*models.KindBasic
	for _, t := range types {
		basic := models.KindBasic{}
		kind, err := basic.FindKindByKindIdentity(t)
		if err != nil {
			return nil, fmt.Errorf("无效的商品类型：%s", t)
		}
		categories = append(categories, kind) // 注意这里的修改，对指针进行解引用
	}
	return categories, nil
}

// GetProductsSimpleInfo 商品信息
// @Summary 获取所有商品的简单信息
// @Description Retrieve a list of all products with basic information and categories.
// @Produce json
// Tags 商品
// @Success 200 {string} ResponseObject
// @Failure 500 {string} ResponseObject
// @Router /products/simple_info [get]
func (CommodityServer) GetProductsSimpleInfo(c *gin.Context) {
	var commodities []models.CommodityBasic
	var count int64

	if err := dao.DB.Preload("Categories").Order("updated_at desc").
		Find(&commodities).
		Count(&count).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取所有的商品信息",
		"data": map[string]interface{}{
			"count":     &count,
			"commodity": commodities,
		},
	})
}

// GetOneProAllInfo godoc
// @Summary 获取特定商品的详细信息
// @Description 根据商品的唯一标识符获取特定商品的详细信息。
// @Produce json
// Tags 商品
// @Param id query string true "商品的唯一标识符" example:"your_commodity_id"
// @Success 200 {string} json {"code": "200", "msg": "成功获取商品信息", "data": CommodityBasic}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /get/one_product_info [post]
func (CommodityServer) GetOneProAllInfo(c *gin.Context) {
	var request struct {
		CommodityIdentity uint `form:"id" binding:"required"`
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  "请求参数错误",
		})
		return
	}
	var count int64
	identity := request.CommodityIdentity
	var commodity models.CommodityBasic
	if err := dao.DB.Preload("Categories", "deleted_at IS NULL").
		Where("commodity_basic.id = ?", identity).
		First(&commodity).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}

	var user models.UserBasic
	if err := dao.DB.Select("avatar,nickname,account,phone_number,score,name,user_identity").
		Where("user_identity = ?", commodity.CommodityIdentity).
		First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取商品所有信息",
		"data": map[string]interface{}{
			"count":     count,
			"commodity": commodity,
			"user":      user,
		},
	})
}

// GetUserAllProList godoc
// @Summary 获取用户所有商品信息
// @Description 获取指定用户的所有商品信息
// @Tags 商品
// @Produce json
// @Param user_identity query string true "用户唯一标识"
// @Success 200 {string} json {"code":200,"msg":"获取商品所有信息","data":{"user_products":{}}}
// @Failure 500 {string} json {"code":500,"msg":"服务器内部错误"}
// @Router /get/user_all_pro_list [post]
func (CommodityServer) GetUserAllProList(c *gin.Context) {
	identity := c.Query("user_identity")
	fmt.Println(identity)
	var count int64
	var user models.UserBasic
	if err := dao.DB.
		Preload("Commodity", func(db *gorm.DB) *gorm.DB {
			return db.Order("updated_at DESC") // 假设CommodityBasic有一个UpdatedAt字段
		}).
		Select("avatar, user_identity, nickname, account, phone_number, score, name, email, wechat_number").
		Where("user_identity = ?", identity).
		First(&user).Count(&count).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}

	// 手动预加载Commodities中的Categories关联数据
	for i := range user.Commodity {
		if err := dao.DB.Model(&user.Commodity[i]).Preload("Categories", func(db *gorm.DB) *gorm.DB {
			return db.Select("name, kind_identity")
		}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取商品所有信息",
		"data": map[string]interface{}{
			"count":         count,
			"user_products": user,
		},
	})
}

// UserModifiesProducts
// @Summary 用户修改商品信息
// @Description 允许用户修改商品详细信息和媒体文件
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param productID query string true "商品ID"
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
// @Router /user/modifies/products [post]
func (CommodityServer) UserModifiesProducts(c *gin.Context) {
	// Parse user information
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
	productID, err := strconv.Atoi(c.Query("productID"))
	fmt.Println(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的商品ID",
		})
		return
	}

	var user models.UserBasic
	if err := dao.DB.Where("account = ?", userClaims.Account).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	var product1 models.CommodityBasic
	if err := dao.DB.Where("id = ?", productID).First(&product1).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品不存在",
		})
		return
	}

	if product1.SoldStatus != 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 400,
			"msg":  "商品信息暂时不能修改",
		})
		return
	}

	// 删除商品与用户的关联关系
	if err := dao.DB.Model(&user).Association("Commodity").Delete(&product1); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// 清除商品与商品类型的多对多关联关系
	if err := dao.DB.Model(&product1).Association("Categories").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}
	// Get form-data values
	types := c.PostFormArray("type")
	title := c.PostForm("title")
	number := c.PostForm("number")
	information := c.PostForm("information")
	price := c.PostForm("price")
	isAuction := c.PostForm("is_auction")
	street := c.PostForm("street")
	city := c.PostForm("city")
	country := c.PostForm("country")
	province := c.PostForm("province")
	contact := c.PostForm("contact")
	postCode := c.PostForm("post_code")

	// Parse number and price
	// Parse number and price
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的参数：商品数量",
		})
		return
	}

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的参数：商品价格",
		})
		return
	}

	// Parse isAuction, handle empty value
	var isAuctionInt int
	if isAuction != "" {
		isAuctionInt, err = strconv.Atoi(isAuction)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "无效的参数：是否拍卖",
			})
			return
		}
	} else {
		// Set a default value or handle it as needed
		isAuctionInt = 0 // Assuming 0 as a default value, update as per your requirement
	}

	// Validate and associate types
	categories, err := validateAndAssociateTypes(types)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	// Create address entry
	modelAddr := models.AddressEntry{
		Street:   street,
		City:     city,
		Country:  country,
		Province: province,
		Contact:  contact,
		PostCode: postCode,
		Identity: pkg.GetAccountNumber(),
	}

	// Find or create user
	exists, existsuser, err := models.UserBasic{}.FindUserByAccount(userClaims.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// Create product
	status := 0
	if isAuctionInt != status {
		status = 1
	}
	addrJson := models.JSONAddress{modelAddr}
	product := &models.CommodityBasic{
		CommodityIdentity: existsuser.UserIdentity,
		Title:             title,
		Number:            numberInt,
		Information:       information,
		Price:             priceFloat,
		SoldStatus:        status,
		Media:             nil,
		IsAuction:         isAuctionInt,
		Address:           addrJson,
		LikeCount:         0,
		CollectCount:      0,
		Categories:        categories,
	}

	// Upload files to OSS
	form, err := c.MultipartForm()
	// Get form-data values
	// ...
	files := form.File["files"] // 注意这里的字段名要与 curl 请求中的一致
	// ...

	for _, file := range files {

		// 保存文件到本地
		localFilePath := "./picture/commodity/" + file.Filename
		fmt.Println("sdf", localFilePath)
		err = c.SaveUploadedFile(file, localFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		objectKey := "picture/commodity" + file.Filename
		err = pkg.UploadAllFile(objectKey, file)
		err = pkg.UploadAllFile(objectKey, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fileURL := fmt.Sprintf("https://%s", objectKey)
		media := models.MediaBasic{Image: fileURL}
		product.Media = append(product.Media, media)
	}

	// Save user and product
	existsuser.Commodity = append(existsuser.Commodity, product)
	if existsuser.SaveUser(existsuser); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}
	if product.CreateCommodity(product); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": "500",
			"msg":  "服务器内部错误",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "成功修改商品信息",
		"data": product,
	})
}

// ClearExistingAssociations 清除商品与商品类型之间的旧关联
func ClearExistingAssociations(productID uint) error {
	// 从数据库中删除与该商品关联的所有旧关联记录
	if err := dao.DB.Where("commodity_basic_id = ?", productID).Delete(&models.KindCommodityRelation{}).Error; err != nil {
		return err
	}

	return nil
}

// UserDeletesProduct
// @Summary 用户删除商品
// @Description 允许用户从其商品列表中删除商品
// @Tags 用户私有方法
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param productID query string true "商品ID"
// @Success 200 {string} json {"code": "200", "msg": "成功删除商品"}
// @Failure 400 {string} json {"code": "400", "msg": "请求无效。服务器无法理解请求"}
// @Failure 401 {string} json {"code": "401", "message": "未授权"}
// @Failure 404 {string} json {"code": "404", "msg": "商品不存在"}
// @Failure 500 {string} json {"code": "500", "msg": "服务器内部错误"}
// @Router /user/deletes/products [post]
func (CommodityServer) UserDeletesProduct(c *gin.Context) {
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

	productID, err := strconv.Atoi(c.Query("productID"))
	fmt.Println(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的商品ID",
		})
		return
	}

	var user models.UserBasic
	if err := dao.DB.Where("account = ?", userClaims.Account).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	var product models.CommodityBasic
	if err := dao.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品不存在",
		})
		return
	}
	// Start a database transaction
	tx := dao.DB.Begin()

	// Check for errors during the transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete the association between user and product
	if err := tx.Model(&user).Association("Commodity").Delete(&product); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// Save the user
	if err := user.SaveUser(&user); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// Clear the association between product and categories
	if err := tx.Model(&product).Association("Categories").Clear(); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// Delete the product itself
	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器内部错误",
		})
		return
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功删除商品",
	})
}
