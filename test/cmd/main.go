package main

import (
	"database/sql"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

type File struct {
	ID         int    `json:"id"`
	FileName   string `json:"file_name"`
	FileURL    string `json:"file_url"`
	UploadTime string `json:"upload_time"`
}

var db *sql.DB
var AccessKeyID = "LTAI5tG17fnfGpThU8nMMjyJ"
var AccessKeySecret = "LruAAgiljPvOrxDaEh2ZYXACugxC2h"

var endpoint = "https://oss-cn-beijing.aliyuncs.com"
var bucketName = "xichen-server"

func init() {
	// 连接数据库
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:33067)/dbname")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	// 连接OSS
	client, err := oss.New(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("Error connecting to OSS:", err)
	}
	backer, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error accessing OSS bucket:", err)
	}
	fmt.Println(backer)
}

func main() {
	defer db.Close()

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MB

	router.GET("/file/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		objectKey := "your-prefix/" + filename

		// 从阿里云OSS下载文件
		url, err := downloadFile(objectKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 返回文件的URL给前端
		c.JSON(http.StatusOK, gin.H{"file_url": url})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 为了演示，将文件保存到本地
		uploadPath := "uploads/"
		if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		filePath := uploadPath + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 上传到阿里云OSS
		objectKey := "your-prefix/" + file.Filename
		err = uploadFile(objectKey, filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 将文件信息保存到数据库
		fileURL := fmt.Sprintf("https://%s/%s", bucketName, objectKey)
		////err = insertFileRecord(file.Filename, fileURL)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return
		//}

		// 返回文件链接或路径给前端
		c.JSON(http.StatusOK, gin.H{"link": fileURL})
	})

	// 启动服务器
	if err := router.Run(":8081"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func insertFileRecord(fileName, fileURL string) error {
	// 将文件信息插入数据库
	stmt, err := db.Prepare("INSERT INTO files (file_name, file_url, upload_time) VALUES (?, ?, NOW())")
	if err != nil {
		return fmt.Errorf("预处理语句时发生错误: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(fileName, fileURL)
	if err != nil {
		return fmt.Errorf("执行语句时发生错误: %v", err)
	}
	return nil
}

func downloadFile(objectKey string) (string, error) {
	client, err := oss.New(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		return "", fmt.Errorf("连接OSS时发生错误: %v", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("访问OSS存储桶时发生错误: %v", err)
	}

	// 生成文件的临时URL，有效期默认为1小时
	url, err := bucket.SignURL(objectKey, oss.HTTPGet, 3600)
	if err != nil {
		return "", fmt.Errorf("生成文件URL时发生错误: %v", err)
	}

	return url, nil
}

func uploadFile(objectKey, filePath string) error {
	client, err := oss.New(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		return fmt.Errorf("连接OSS时发生错误: %v", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("访问OSS存储桶时发生错误: %v", err)
	}

	// 上传文件到OSS
	err = bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		return fmt.Errorf("上传文件到OSS时发生错误: %v", err)
	}

	return nil
}
