package alicloud

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-playground/validator/v10"
	"github.com/yyg192/Go_FileStation/store"
)

func NewUploader(endpoint, accessKeyId, accessKeySecret string) (store.Uploader, error) {
	uploader := &alicloud{
		Endpoint:        endpoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		listener:        NewListener(),
	}
	if err := uploader.validate(); err != nil {
		return nil, err
	}
	return uploader, nil
}

type alicloud struct {
	Endpoint        string `validator:"required"`
	AccessKeyId     string `validator:"required"`
	AccessKeySecret string `validator:"required"`
	listener        oss.ProgressListener
}

var validate = validator.New()

func (a *alicloud) validate() error {
	return validate.Struct(a)
	//校验结构体a中的字段值是否合法（是否满足required，即是否被赋值了）
}

func (a *alicloud) UploadFile(bucketName, objectKey, filePath string) error {
	if filePath == "" || objectKey == "" {
		return fmt.Errorf("filePath missed")
	}
	client, err := oss.New(a.Endpoint, a.AccessKeyId, a.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(filePath, filePath, oss.Progress(a.listener))
	//第一个参数是上传到oss里面的文件路径，第二个参数是文件的本地路径
	if err != nil {
		return err
	}

	//打印下载URL
	//1. 获取 object对象
	signedUrl, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return fmt.Errorf("sign file down load url error %s", err)
	}
	fmt.Printf("下载链接: %s\n", signedUrl)
	fmt.Println("\n注意:文件下载有效期为1天，中转站保存时间为3天，请及时下载")
	return nil
}
