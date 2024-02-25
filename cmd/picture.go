package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var AccessKeyID = "LTAI5tQdpsht36XJqYZevPGn"
var AccessKeySecret = "dyThns69pEwBcGl1AjBby1YsGEyDcm"

var endpoint = "https://oss-cn-beijing.aliyuncs.com"
var bucketName = "xichen-server"

func main() {

	// 创建 OSS 客户端
	client, err := oss.New(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取存储桶
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	objectKey := "9fd8e447-0250-4741-8c8a-18b19e272a91.png"
	// 生成图片对象的访问 URL
	url, err := bucket.SignURL(objectKey, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Image URL:", url)
}
