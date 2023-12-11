package dao

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
)

var RDB *redis.Client

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewRDB() {
	// 设置 Viper 的配置文件名和路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)

	}

	// 使用 Viper 获取 MySQL 配置
	var redisConfig RedisConfig
	if err := viper.UnmarshalKey("redis", &redisConfig); err != nil {
		fmt.Println("Error unmarshalling Redis config:", err)
	}

	// 打印配置信息
	fmt.Printf("Redis Config: %+v\n", redisConfig)
	// 初始化 redis 连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + strconv.Itoa(redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	RDB = rdb
}
