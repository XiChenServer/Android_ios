package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = newDB()

type MySQLConfig struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Database  string
	Charset   string
	ParseTime bool
}

func newDB() *gorm.DB {
	// 设置 Viper 的配置文件名和路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)

	}

	// 使用 Viper 获取 MySQL 配置
	var mysqlConfig MySQLConfig
	if err := viper.UnmarshalKey("mysql", &mysqlConfig); err != nil {
		fmt.Println("Error unmarshalling MySQL config:", err)
	}

	// 打印配置信息
	fmt.Printf("MySQL Config: %+v\n", mysqlConfig)
	// 初始化 Gorm 连接
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
		mysqlConfig.Charset,
		mysqlConfig.ParseTime)))

	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	return db
}
