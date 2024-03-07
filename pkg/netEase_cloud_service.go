package pkg

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
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

func SendEmailVerificationCode1(recipientEmail, verificationCode string) error {
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
		return err
	}

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
		return err
	}

	fmt.Println("Email sent successfully.")
	return nil
}

func SendEmailVerificationCode(email, verificationCode1 string) error {
	////读取配置文件
	//configData, err := ioutil.ReadFile("config/config.yaml")
	//if err != nil {
	//    log.Fatal("无法读取配置文件:", err)
	//}
	//
	//// 解析配置文件
	//var config Config
	//err = yaml.Unmarshal(configData, &config)
	//if err != nil {
	//    log.Fatal("无法解析配置文件:", err)
	//}
	//发送对象
	recipient := email
	// 生成验证码
	verificationCode := verificationCode1

	// 构建邮件内容
	subject := "验证码"
	body := fmt.Sprintf("你的验证码是：%s", verificationCode)
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPortStr := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	// 创建邮件消息
	message := gomail.NewMessage()
	message.SetHeader("From", "15294440097@163.com")
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)
	// 转换 SMTP_PORT 为整数
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		fmt.Println("Error converting SMTP_PORT to integer:", err)
		return err
	}
	// 创建SMTP客户端
	dialer := gomail.NewDialer(smtpServer, smtpPort, senderEmail, senderPassword)

	// 发送邮件
	err = dialer.DialAndSend(message)
	if err != nil {
		fmt.Println("发送邮件失败:", err)
		return err
	}

	return nil
}

//func SendCode(email string) string {
//	////读取配置文件
//	//configData, err := ioutil.ReadFile("config/config.yaml")
//	//if err != nil {
//	//    log.Fatal("无法读取配置文件:", err)
//	//}
//	//
//	//// 解析配置文件
//	//var config Config
//	//err = yaml.Unmarshal(configData, &config)
//	//if err != nil {
//	//    log.Fatal("无法解析配置文件:", err)
//	//}
//	//发送对象
//	recipient := email
//	// 生成验证码
//	verificationCode := generateVerificationCode()
//
//	// 构建邮件内容
//	subject := "验证码"
//	body := fmt.Sprintf("你的验证码是：%s", verificationCode)
//
//	// 创建邮件消息
//	message := gomail.NewMessage()
//	message.SetHeader("From", "19891294013@163.com")
//	message.SetHeader("To", recipient)
//	message.SetHeader("Subject", subject)
//	message.SetBody("text/plain", body)
//
//	// 创建SMTP客户端
//	dialer := gomail.NewDialer("smtp.163.com", 465, "19891294013@163.com", "DRCJMYFWIGGKGSWM")
//
//	// 发送邮件
//	err := dialer.DialAndSend(message)
//	if err != nil {
//		fmt.Println("发送邮件失败:", err)
//		return ""
//	}
//
//	return verificationCode
//}
//
//func generateVerificationCode() string {
//	// 设置随机数种子
//	rand.Seed(time.Now().UnixNano())
//
//	// 生成6位验证码
//	code := rand.Intn(899999) + 100000
//
//	// 将验证码转换为字符串
//	codeStr := strconv.Itoa(code)
//
//	return codeStr
//
//}
