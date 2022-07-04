package ecs

import (
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/Dao/providers/host"
)

// 这个名字感觉取的不是很好 大概就是说获取一页的ECS数据
func (eh *EcsObtainer) M_GetEcsPageInfo(req *sdk_ecs.DescribeInstancesRequest) (*host.HostSet, error) {
	hostSet := host.CreateHostSet()
	resp, err := eh.m_ecs_client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}
	hostSet.M_total = int32(resp.TotalCount)
	hostSet.M_items = eh.m_TransferHostSet(resp.Instances.Instance).M_items
	return hostSet, nil
}
