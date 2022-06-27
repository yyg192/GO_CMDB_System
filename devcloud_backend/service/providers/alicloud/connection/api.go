package connection

import "fmt"

func CreateAliCloudClient(regionId string, accessKey string, accessSecret string) *AliCloudClient {
	return &AliCloudClient{
		m_regionId:     regionId,
		m_accessKey:    accessKey,
		m_accessSecret: accessSecret,
	}
}

func PrintAliCloudClientInfo(acc *AliCloudClient) {
	fmt.Printf("region_id: %s\n access_key: %s\n access_secret: %s\n",
		acc.m_regionId, acc.m_accessKey, acc.m_accessSecret)
}
