package connection

import (
	"fmt"

	sdk_bss "github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	sdk_rds "github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	sdk_sts "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	sdk_oss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func (acc *AliCloudClient) M_EcsClientConnection() (*sdk_ecs.Client, error) {
	if acc.M_ecsClient != nil {
		return acc.M_ecsClient, nil
	}
	ecs_client, err := sdk_ecs.NewClientWithAccessKey(acc.m_regionId, acc.m_accessKey, acc.m_accessSecret)
	if err != nil {
		return nil, err
	}
	acc.M_ecsClient = ecs_client
	return ecs_client, nil
}

func (acc *AliCloudClient) M_RdsClientConnection() (*sdk_rds.Client, error) {
	return nil, fmt.Errorf("function RdsClientConnection not finished yet")
}

func (acc *AliCloudClient) M_BssClientConnection() (*sdk_bss.Client, error) {
	return nil, fmt.Errorf("function BssClientConnection not finished yet")
}

func (acc *AliCloudClient) M_OssClientConnection() (*sdk_oss.Client, error) {
	return nil, fmt.Errorf("function OssClientConnection not finished yet")
}

func (acc *AliCloudClient) M_Account() (string, error) {
	// 通过sdk拿到云商该账户的accountId
	req := sdk_sts.CreateGetCallerIdentityRequest()
	stsClient, err := sdk_sts.NewClientWithAccessKey(acc.m_regionId, acc.m_accessKey, acc.m_accessSecret)
	stsClient.GetConfig().WithScheme("HTTPS")
	if err != nil {
		return "", fmt.Errorf("unable to initialize the STS client: %#v", err)
	}
	/*
		阿里云STS（Security Token Service）是阿里云提供的一种临时访问权限管理服务。
		RAM提供RAM用户和RAM角色两种身份。其中，RAM角色不具备永久身份凭证，
		而只能通过STS获取可以自定义时效和访问权限的临时身份凭证，
		即安全令牌（STS Token）。
	*/
	stsClient.AppendUserAgent("XXX", "1.0")
	identity, err := stsClient.GetCallerIdentity(req)
	if err != nil {
		return "", err
	}
	if identity == nil {
		return "", fmt.Errorf("caller identity not found")
	}
	return identity.AccountId, nil
}
