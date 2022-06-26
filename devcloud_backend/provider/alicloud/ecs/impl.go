package ecs

import (
	"time"

	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/entity/host"
)

/**
开发日志：
2022.6.26 Tags没完成，sdk的用法还不是很熟悉，先搁置一下，但是别忘了！！
**/

/**
这个包专门用来进行ecs操作
**/
type EcsHandler struct {
	m_ecs_client *sdk_ecs.Client
	*host.AccountGetter
}

func CreateEcsHandler(client *sdk_ecs.Client) *EcsHandler {
	return &EcsHandler{
		m_ecs_client:  client,
		AccountGetter: &host.AccountGetter{},
	}
}
func (eh *EcsHandler) GetInstancesDescription(req *sdk_ecs.DescribeInstancesRequest) (*host.HostSet, error) {
	host_set := host.NewHostSet()
	resp, err := eh.m_ecs_client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}
	host_set.Total = int(resp.TotalCount)
	host_set.Items = eh.PlaceToHostSet(resp.Instances.Instance).Items

	return host_set, nil
}

func (eh *EcsHandler) PlaceToHostSet(unfilter_items []sdk_ecs.Instance) *host.HostSet {
	// 云商会传来一堆描述信息字段(unfilter_items)，因为我们只需要其中的一部分
	// 并把所需的那一部分填进HostSet中返回
	host_set := host.NewHostSet()
	for i := range unfilter_items {
		host_set.Add(eh.PlaceToHost(unfilter_items[i]))
	}
	return host_set
}

func (eh *EcsHandler) PlaceToHost(ins sdk_ecs.Instance) *host.Host {
	host := host.NewDefaultHost()
	host.Base.Id = ins.InstanceId
	host.Base.Vendor = 0 //这里一定要注意！我暂时给个0而已，后面还要统一起来的！
	host.Base.Region = ins.RegionId
	host.Base.Zone = ins.ZoneId
	host.Base.CreateAt = eh.parseTime(ins.CreationTime)
	//host.Base.ResourceHash = ins.
	//host.Base.DescribeHash = ins.

	host.Resource.ExpireAt = eh.parseTime(ins.ExpiredTime)
	//host.Resource.Category    =
	host.Resource.Type = ins.InstanceType
	host.Resource.Name = ins.InstanceName
	host.Resource.Description = ins.Description
	host.Resource.Status = ins.Status
	/* host.Resource.Tags = eh.transferTags(ins.Tags.Tag) 暂时先不处理这个*/
	//host.Resource.UpdateAt    =
	host.Resource.SyncAccount = eh.GetAccountId()
	host.Resource.PublicIP = ins.PublicIpAddress.IpAddress
	host.Resource.PrivateIP = eh.parsePrivateIp(ins)
	host.Resource.PayType = ins.InstanceChargeType

	host.Describe.CPU = int(ins.CPU)
	host.Describe.Memory = int(ins.Memory)
	host.Describe.GPUAmount = int(ins.GPUAmount)
	host.Describe.GPUSpec = ins.GPUSpec
	host.Describe.OSType = ins.OsType
	host.Describe.OSName = ins.OSName
	host.Describe.SerialNumber = ins.SerialNumber
	host.Describe.ImageID = ins.ImageId
	host.Describe.InternetMaxBandwidthOut = int(ins.InternetMaxBandwidthOut)
	host.Describe.InternetMaxBandwidthIn = int(ins.InternetMaxBandwidthIn)
	host.Describe.KeyPairName = []string{ins.KeyPairName}
	host.Describe.SecurityGroups = ins.SecurityGroupIds.SecurityGroupId
	return host
}

func (eh *EcsHandler) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04Z", t)
	if err != nil {
		//eh.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

func (eh *EcsHandler) parsePrivateIp(ins sdk_ecs.Instance) []string {
	ips := []string{} //空切片
	//优先查主网卡的私网IP地址
	/**
		"NetworkInterfaces" : {
	        "NetworkInterface" : [ {
	          "Type" : "Primary",
	          "PrimaryIpAddress" : "172.17.**.***",
	          "NetworkInterfaceId" : "eni-2zeh9atclduxvf1z****",
	          "MacAddress" : "00:16:3e:32:b4:**",
	          "PrivateIpSets" : {
	            "PrivateIpSet" : [ {
	              "PrivateIpAddress" : "172.17.**.**",
	              "Primary" : true
	            } ]
	          }
	        } ]
		},
		**/
	for _, ni := range ins.NetworkInterfaces.NetworkInterface {
		for _, ip := range ni.PrivateIpSets.PrivateIpSet {
			if ip.Primary {
				ips = append(ips, ip.PrivateIpAddress)
			}
		}
	}

	if len(ips) > 0 {
		return ips
	}
	/**
	经典内网
	"InnerIpAddress": {
		"IpAddress": [
		"10.xx.xx.xx"
		]
	},
	**/
	if len(ins.InnerIpAddress.IpAddress) > 0 {
		return ins.InnerIpAddress.IpAddress
	}

	/**
	如果不是经典内网就是Vpc网络
	**/
	return ins.VpcAttributes.PrivateIpAddress.IpAddress
}

// func (eh *EcsHandler) transferTags(t []sdk_ecs.Tag) error {
// 	return fmt.Errorf("function transferTags is not finished yet")
// }
