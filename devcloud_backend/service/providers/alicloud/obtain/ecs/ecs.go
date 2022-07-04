package ecs

import (
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/Dao/providers/host"
)

type EcsObtainer struct {
	m_ecs_client *sdk_ecs.Client
	*host.AccountGetter
	//resource.AccountGetter
}

func CreateEcsObtainer(client *sdk_ecs.Client) *EcsObtainer {
	return &EcsObtainer{
		m_ecs_client:  client,
		AccountGetter: &host.AccountGetter{},
	}
}
