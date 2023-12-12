package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

func init() {
	// 加载 .env 文件中的环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// 网易SMTP服务器地址和端口
	// 使用 os.Getenv 获取环境变量的值
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPortStr := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	fmt.Println("SMTP Server:", smtpServer)
	fmt.Println("SMTP Port:", smtpPortStr)

	// 转换 SMTP_PORT 为整数
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		fmt.Println("Error converting SMTP_PORT to integer:", err)
		return
	}

	// 收件人邮箱地址
	recipientEmail := "3551906947@qq.com"

	// 生成验证码
	verificationCode := "123456" // 请替换为您生成的实际验证码

	// 构建邮件主体
	emailBody := fmt.Sprintf(`
		<html>
			<body>
				<p>尊敬的用户，您的验证码是：%s</p>
				<p>请在有效期内使用验证码完成相关操作。</p>
			</body>
		</html>
	`, verificationCode)

	message := []byte(fmt.Sprintf("To: %s\r\n", recipientEmail) +
		"Subject: 验证码邮件\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		emailBody)

	// 连接到网易SMTP服务器
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	// 发送邮件
	err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully.")
}
