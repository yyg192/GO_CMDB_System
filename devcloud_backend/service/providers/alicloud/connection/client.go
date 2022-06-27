package connection

import (
	sdk_bss "github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	sdk_rds "github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

type AliCloudClient struct {
	/** private **/
	m_regionId     string
	m_accessKey    string
	m_accessSecret string
	/** public **/
	M_ecsClient *sdk_ecs.Client
	M_rdsClient *sdk_rds.Client
	M_bssClient *sdk_bss.Client
}
