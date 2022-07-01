package ecs

import (
	"time"

	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/Dao/providers/host"
)

/**
2. query.go负责利用阿里云提供的sdk，从第三方拉去实例描述信息，
而transfer.go则负责将第三方传递过来的描述信息转为本地数据对象（选取我们需要的描述信息字段存进我们的结构体对象）
**/

func (eh *EcsHandler) m_TransferHostSet(unfilter_items []sdk_ecs.Instance) *host.HostSet {
	// 云商会传来一堆描述信息字段(unfilter_items)，因为我们只需要其中的一部分
	// 并把所需的那一部分填进HostSet中返回
	host_set := host.CreateHostSet()
	for i := range unfilter_items {
		host_set.M_Add(eh.m_TransferHost(unfilter_items[i]))
	}
	return host_set
}

func (eh *EcsHandler) m_TransferHost(ins sdk_ecs.Instance) *host.Host {
	host := host.CreateHost()
	host.BasicInformation.Id = ins.InstanceId
	host.BasicInformation.Vendor = 0 //这里一定要注意！我暂时给个0而已，后面还要统一起来的！
	host.BasicInformation.Region = ins.RegionId
	host.BasicInformation.Zone = ins.ZoneId
	host.BasicInformation.CreateAt = eh.m_ParseTime(ins.CreationTime)
	//host.Base.ResourceHash = ins.
	//host.Base.DescribeHash = ins.

	host.ResourceInformation.ExpireAt = eh.m_ParseTime(ins.ExpiredTime)
	//host.Resource.Category    =
	host.ResourceInformation.Type = ins.InstanceType
	host.ResourceInformation.Name = ins.InstanceName
	host.ResourceInformation.Description = ins.Description
	host.ResourceInformation.Status = ins.Status
	/* host.Resource.Tags = eh.transferTags(ins.Tags.Tag) 暂时先不处理这个*/
	//host.Resource.UpdateAt    =
	host.ResourceInformation.SyncAccount = eh.GetAccountId()
	host.ResourceInformation.PublicIP = ins.PublicIpAddress.IpAddress
	host.ResourceInformation.PrivateIP = eh.m_ParsePrivateIp(ins)
	host.ResourceInformation.PayType = ins.InstanceChargeType

	host.DescribeInformation.CPU = int(ins.CPU)
	host.DescribeInformation.Memory = int(ins.Memory)
	host.DescribeInformation.GPUAmount = int(ins.GPUAmount)
	host.DescribeInformation.GPUSpec = ins.GPUSpec
	host.DescribeInformation.OSType = ins.OsType
	host.DescribeInformation.OSName = ins.OSName
	host.DescribeInformation.SerialNumber = ins.SerialNumber
	host.DescribeInformation.ImageID = ins.ImageId
	host.DescribeInformation.InternetMaxBandwidthOut = int(ins.InternetMaxBandwidthOut)
	host.DescribeInformation.InternetMaxBandwidthIn = int(ins.InternetMaxBandwidthIn)
	host.DescribeInformation.KeyPairName = []string{ins.KeyPairName}
	host.DescribeInformation.SecurityGroups = ins.SecurityGroupIds.SecurityGroupId
	return host
}

func (eh *EcsHandler) m_ParseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04Z", t)
	if err != nil {
		//eh.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

func (eh *EcsHandler) m_ParsePrivateIp(ins sdk_ecs.Instance) []string {
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

// func (eh *EcsHandler) m_TransferTags(t []sdk_ecs.Tag) error {
// 	return fmt.Errorf("function transferTags is not finished yet")
// }
