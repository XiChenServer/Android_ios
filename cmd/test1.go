package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
	// 从环境变量中获取访问凭证。运行本代码示例之前，请先配置环境变量。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// yourEndpoint填写Bucket对应的自定义域名。
	// oss.UseCname(true)用于开启CNAME。CNAME用于将自定义域名绑定至存储空间。
	client, err := oss.New("yourEndpoint", "", "", oss.SetCredentialsProvider(&provider), oss.UseCname(true))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Printf("client:%#v\n", client)
}
