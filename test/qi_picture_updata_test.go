package test

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
	"testing"
)

func Test_qi_picture(T *testing.T) {
	accessKey := "y_XTiaH5dywx_R-J-twejWCQRXvBd5jI54YT9ihT"
	secretKey := "2g0S7zGWZ_zca0BVwYTeugUoZJepYLsYjd5bKGir"
	bucket := "taoniuma"
	// 构建一个七牛云存储的配置
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 这里选择华南地区
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	// 生成上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	// 生成mac
	mac := qbox.NewMac(accessKey, secretKey)
	// 构建表单上传的对象
	uploader := storage.NewFormUploader(&cfg)
	upToken := putPolicy.UploadToken(mac)
	// 要上传文件的本地路径
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
	ret := storage.PutRet{}
	// 传递一个有效的 context
	ctx := context.Background()
	// 上传文件，传递 context 参数
	putExtra := storage.PutExtra{}
	err = uploader.PutWithoutKey(ctx, &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println("上传失败:", err)
		return
	}
	fmt.Println("上传成功，图片URL为:", "http://s9isqyrv9.hn-bkt.clouddn.com/"+ret.Key)
}
