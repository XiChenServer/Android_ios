package pkg

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"strings"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
var AccessKeyID = "LTAI5tQdpsht36XJqYZevPGn"
var AccessKeySecret = "dyThns69pEwBcGl1AjBby1YsGEyDcm"

var endpoint = "https://oss-cn-beijing.aliyuncs.com"
var bucketName = "xichen-server"

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendPhoneCode(phone, code string) error {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String(AccessKeyID), tea.String(AccessKeySecret))
	if _err != nil {
		return _err
	}
	//codeStr := "{\"code\":\"" + code + "\"}"
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("XIChenServer"),
		TemplateCode:  tea.String("SMS_464065886"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return nil
}

// UploadAvatarFromForm 从表单文件上传到阿里云 OSS
func UploadAvatarFromForm(fileName string, file io.Reader) error {
	// 打开本地文件

	// 创建阿里云 OSS 客户端
	client, err := oss.New(endpoint, AccessKeyID, AccessKeySecret)
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
	return err
}

// DownAvatarFromOSS 从阿里云OSS获取头像
func DownAvatarFromOSS(fileName string) (io.ReadCloser, error) {
	// 创建阿里云 OSS 客户端
	client, err := oss.New(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	// 下载文件
	fileReader, err := bucket.GetObject(fileName)
	if err != nil {
		return nil, err
	}

	return fileReader, nil
}
