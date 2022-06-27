package ecs

import (
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/model/host"
)

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
