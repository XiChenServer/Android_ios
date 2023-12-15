package test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func uploadHandler(c *gin.Context) {
	// 解析多部分表单数据
	err := c.Request.ParseMultipartForm(10 << 20) // 限制最大文件大小为 10MB
	if err != nil {
		c.JSON(400, gin.H{"error": "Error parsing multipart form"})
		return
	}

	// 获取所有上传的文件
	form := c.Request.MultipartForm
	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(400, gin.H{"error": "No files provided"})
		return
	}

	// 处理每个上传的文件
	for _, file := range files {
		// 打印文件信息
		fmt.Printf("Uploaded File: %s, Size: %d bytes\n", file.Filename, file.Size)
		suffix := getFileExtensionWithoutDot(file.Filename)
		fmt.Println(suffix)
		// 直接获取文件内容的 io.Reader
		fileReader, err := file.Open()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error opening file"})
			return
		}
		defer fileReader.Close()

		// 在这里你可以使用 fileReader 进行进一步的操作
		// 例如，你可以将其传递给某个需要 io.Reader 的函数
		// DoSomethingWithReader(fileReader)
		image("user2_test."+suffix, fileReader)
	}

	c.JSON(200, gin.H{"message": "Files uploaded successfully"})
}

func Test_imageUpload(t *testing.T) {
	r := gin.Default()
	r.POST("/upload", uploadHandler)

	// 创建用于保存上传文件的目录
	os.Mkdir("uploads", os.ModePerm)

	// 启动 HTTP 服务器，监听在 :8080 端口上
	r.Run(":8080")
}
func image(fileName string, file io.Reader) {
	// 替换为你的阿里云 OSS 访问密钥信息
	accessKeyID := "LTAI5tQdpsht36XJqYZevPGn"
	accessKeySecret := "dyThns69pEwBcGl1AjBby1YsGEyDcm"
	endpoint := "https://oss-cn-beijing.aliyuncs.com"
	bucketName := "xichen-server"

	// 打开本地文件

	// 创建阿里云 OSS 客户端
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}

	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}

	// 上传文件到 OSS
	err = bucket.PutObject(fileName, file)
	if err != nil {
		panic(err)
	}
}
func getFileExtensionWithoutDot(filename string) string {
	ext := filepath.Ext(filename)
	return strings.TrimPrefix(ext, ".")
}
