package ecs

import (
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/Dao/providers/host"
)

func (eh *EcsHandler) M_GetEcsHostGroupFromAlicloud(req *sdk_ecs.DescribeInstancesRequest) (*host.HostSet, error) {
	host_set := host.CreateHostSet()
	resp, err := eh.m_ecs_client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}
	host_set.M_total = int32(resp.TotalCount)
	host_set.M_items = eh.m_TransferHostSet(resp.Instances.Instance).M_items
	return host_set, nil
}
