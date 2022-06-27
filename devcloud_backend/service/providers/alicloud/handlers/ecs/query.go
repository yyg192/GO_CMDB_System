package ecs

import (
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/model/host"
)

func (eh *EcsHandler) M_GetAllInstancesDescriptionFromAlicloud(req *sdk_ecs.DescribeInstancesRequest) (*host.HostSet, error) {
	host_set := host.NewHostSet()
	resp, err := eh.m_ecs_client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}
	host_set.Total = int32(resp.TotalCount)
	host_set.Items = eh.m_TransferHostSet(resp.Instances.Instance).Items

	return host_set, nil
}
