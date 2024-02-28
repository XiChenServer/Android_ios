package main

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"os"
)

const (
	AccessKey = "y_XTiaH5dywx_R-J-twejWCQRXvBd5jI54YT9ihT"
	SerectKey = "2g0S7zGWZ_zca0BVwYTeugUoZJepYLsYjd5bKGir"
	Bucket    = "taoniuma"
	ImgUrl    = "http://s9isqyrv9.hn-bkt.clouddn.com/"
)

func main() {

	localFile := "/home/zwm/GolandProjects/Android_ios/uploads/截图 2023-11-22 21-45-21.png"
	// 初始化一个进度记录对象
	// 打开本地文件
	file, err := os.Open(localFile)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败:", err)
		return
	}
	fileSize := fileInfo.Size()

	UploadToQiNiu(file, fileSize)
}

// 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(file multipart.File, fileSize int64) (int, string) {
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {

		return 0, err.Error()
	}
	url := ImgUrl + ret.Key
	fmt.Println(url)
	return 200, url
}
