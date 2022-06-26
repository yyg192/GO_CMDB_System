package connection

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func (acc *AliCloudClient) EcsClientConnection() (*sdk_ecs.Client, error) {
	if acc.M_ecs_client != nil {
		return acc.M_ecs_client, nil
	}
	ecs_client, err := sdk_ecs.NewClientWithAccessKey(acc.m_region_id, acc.m_access_key, acc.m_access_secret)
	if err != nil {
		return nil, err
	}
	acc.M_ecs_client = ecs_client
	return ecs_client, nil
}

func (acc *AliCloudClient) RdsClientConnection() (*rds.Client, error) {
	return nil, fmt.Errorf("function RdsClientConnection not finished yet")
}

func (acc *AliCloudClient) BssClientConnection() (*bssopenapi.Client, error) {
	return nil, fmt.Errorf("function BssClientConnection not finished yet")
}

func (acc *AliCloudClient) OssClientConnection() (*oss.Client, error) {
	return nil, fmt.Errorf("function OssClientConnection not finished yet")
}

// func (acc *AliCloudClient) GetAccount() (string, error) {

// }
