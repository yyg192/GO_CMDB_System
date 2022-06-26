package connection

import (
	"fmt"

	alisdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

var (
	Test_region_id     string = "cn-hangzhou"
	Test_access_key    string = "LTAI5tQFjrUGN8KsZmCRrzfP"
	Test_access_secret string = "nSHMjM7fWYubDc4Awg5MzjCPTrZaq6"
)

type AliCloudClient struct {
	/** private **/
	m_region_id     string
	m_access_key    string
	m_access_secret string
	/** public **/
	M_ecs_client *alisdk_ecs.Client
}

func CreateAliCloudClient(region_id string, access_key string, access_secret string) *AliCloudClient {
	return &AliCloudClient{
		m_region_id:     region_id,
		m_access_key:    access_key,
		m_access_secret: access_secret,
	}

}

func (acc *AliCloudClient) PrintAliCloudClientInfo() {
	fmt.Printf("region_id: %s\n access_key: %s\n access_secret: %s\n",
		acc.m_region_id, acc.m_access_key, acc.m_access_secret)
}
