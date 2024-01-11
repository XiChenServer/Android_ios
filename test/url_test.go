package test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"testing"
)

var AccessKeyID = "LTAI5tQdpsht36XJqYZevPGn"
var AccessKeySecret = "dyThns69pEwBcGl1AjBby1YsGEyDcm"

var endpoint = "xichen.love"
var bucketName = "xichen-server"

func Test_url(t *testing.T) {
	// 创建OSS客户端
	client, err := oss.New(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("Error creating OSS client:", err)
		return
	}

	// 获取存储桶
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error obtaining bucket:", err)
		return
	}

	// 设置文件在 OSS 中的存储路径
	ossStoragePath := "test/test.txt"

	// 生成签名 URL，有效期设置为1小时
	signedURL, err := bucket.SignURL(ossStoragePath, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println("Error generating signed URL:", err)
		return
	}

	fmt.Println("Signed URL:", signedURL)
}
