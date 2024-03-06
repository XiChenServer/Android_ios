package main

import (
	"fmt"
	"os"
)

func main() {
	// 设置名为 "MY_SECRET_KEY" 的环境变量
	err := os.Setenv("AccessKeyID", "LTAI5tG17fnfGpThU8nMMjyJ")
	if err != nil {
		fmt.Println("无法设置环境变量:", err)
		return
	}

	err = os.Setenv("AccessKeySecret", "LruAAgiljPvOrxDaEh2ZYXACugxC2h")
	if err != nil {
		fmt.Println("无法设置环境变量:", err)
		return
	}
	// 现在可以在程序的其他部分读取此环境变量
	secretKey := os.Getenv("AccessKeySecret")
	fmt.Println("已设置的秘钥:", secretKey)
}
