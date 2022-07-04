package cloudHandler

import (
	"context"

	"github.com/yyg192/GO_CMDB_System/api/providers/pager"
)

type AbstractEcsHandler interface {
	M_GetHostGroup(req *QueryEcsGroupRequest) pager.AbstractPager          //从云商获取一定数量的Ecs主机实例，这个数量就是一个page的大小
	M_GetHostDetailOne(ctx context.Context, req *QueryEcsDetailOneRequest) //从云商获取一个Ecs主机实例的详细信息
}

type QueryEcsGroupRequest struct {
	// 查询一组Ecs实例描述的请求
	M_Rate float64 `json:"rate"`
}
type QueryEcsDetailOneRequest struct {
	// 查询一个Ecs实例的详细描述信息的请求
	M_Id string `json:"id"`
}

func CreateQueryEcsGroupRequest() *QueryEcsGroupRequest {
	return &QueryEcsGroupRequest{
		M_Rate: 5,
	}
}

func CreateQueryEcsRequestGroupWithRate(rate int32) *QueryEcsGroupRequest {
	return &QueryEcsGroupRequest{
		M_Rate: float64(rate),
	}
}
