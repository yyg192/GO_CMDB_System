package cmd

import (
	"fmt"
	"path"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/yyg192/Go_FileStation/store"
	"github.com/yyg192/Go_FileStation/store/provider/alicloud"
)

const (
	defaultBucketName = "cloud-station-yyg"
	defaultEndPoint   = "oss-cn-hangzhou.aliyuncs.com"
)

var (
	bucketName     string
	uploadFilePath string
	bucketEndPoint string
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件到中转站",
	Long:  `上传文件到中转站`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getUploader()
		if err != nil {
			return err
		}
		if uploadFilePath == "" {
			return fmt.Errorf("upload file path is missing")
		}
		day := time.Now().Format("20060102")
		fn := path.Base(uploadFilePath) //Base return the last element of the path
		ok := fmt.Sprintf("%s/%s", day, fn)
		err = p.UploadFile(bucketName, ok, uploadFilePath) //上传到oss的文件名为  日期/文件名
		if err != nil {
			return err
		}
		return nil
	},
}

func getUploader() (store.Uploader, error) {
	switch ossProvider {
	case "aliCloud":
		AKPrompt := &survey.Password{
			Message: "Please Input Ali Access Id",
		}
		survey.AskOne(AKPrompt, &aliAccessID)
		SKPrompt := &survey.Password{
			Message: "Please input Ali Access Secret Key",
		}
		survey.AskOne(SKPrompt, &aliAccessKey)

		return alicloud.NewUploader(bucketEndPoint, aliAccessID, aliAccessKey)
	case "tencentCloud":
		return nil, fmt.Errorf("not impl")
	case "minio":
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknown uploader %s", ossProvider)
	}
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadFilePath, "file_path", "f", "", "upload file path")
	uploadCmd.PersistentFlags().StringVarP(&bucketName, "bucket_name", "b", defaultBucketName, "upload oss bucket name")
	uploadCmd.PersistentFlags().StringVarP(&bucketEndPoint, "bucket_endpoint", "e", defaultEndPoint, "upload oss endpoint")
	RootCmd.AddCommand(uploadCmd)
	//将uploadCmd作为RootCmd的嵌套
}
