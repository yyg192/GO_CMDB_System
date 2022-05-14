package alicloud_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yyg192/Go_FileStation/store/provider/alicloud"
)

//测试包编译构建的时候是不会被打包进去的

var (
	endpoint        = "oss-cn-hangzhou.aliyuncs.com"
	accessKeyId     = "LTAI5tFJsRfoQc4tHyycDydr"
	accessKeySecret = "jLpRuhQgxJzgJvMFw86yiQhAn1Fsky"
	bucketName      = "cloud-station-yyg"
	filePath        = "store.go"
)

func TestUploadFile(t *testing.T) {
	should := assert.New(t)
	uploader, err := alicloud.NewUploader(endpoint, accessKeyId, accessKeySecret)
	if should.NoError(err) {
		err = uploader.UploadFile(bucketName, filePath, filePath)
		should.NoError(err)
	}
}
