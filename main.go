package main

import (
	_ "Android_ios/docs"
	"Android_ios/models"
	"Android_ios/router"
)

// @title 淘牛马
// @version 1.0
// @description 用于家畜交易
// @termsOfService http://swagger.io/terms/
// @contact.name phone_number 15294440097
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:13000
// @BasePath
func main() {
	models.Init()
	r := router.Router()
	r.Run(":13000")
}
