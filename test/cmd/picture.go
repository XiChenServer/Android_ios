package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func main() {
	/// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	//provider, err := oss.NewEnvironmentVariableCredentialsProvider()

	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("https://oss-cn-beijing.aliyuncs.com", "LTAI5tQdpsht36XJqYZevPGn", "dyThns69pEwBcGl1AjBby1YsGEyDcm")
	if err != nil {
		HandleError(err)
	}

	// 指定图片所在Bucket的名称，例如examplebucket。
	bucketName := "xichen-server"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	// 指定图片名称。如果图片不在Bucket根目录，需携带文件完整路径，例如exampledir/example.jpg。
	ossImageName := "your-prefix/9682B802FF091A4746DCC98526E2FE8B.jpg"
	// 生成带签名的URL，并指定过期时间为600s。
	signedURL, err := bucket.SignURL(ossImageName, oss.HTTPGet, 600, oss.Process("image/format,png"))
	if err != nil {
		HandleError(err)
	} else {
		fmt.Println(signedURL)
	}
}
