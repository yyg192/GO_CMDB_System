package ecs

import (
	"context"

	sdk_req "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/yyg192/GO_CMDB_System/api/providers/common/pager"
	"github.com/yyg192/GO_CMDB_System/api/providers/common/set"
)

/**
页迭代器，专门用来以页为单位获取和迭代主机实例资源信息
EcsPager只需要继承BasePager，然后实现获取当前页的资源的函数，并且通过M_RegistFuncGetCurrentPageData
进行注册就好了。
**/

type EcsPager struct {
	/**
	# Implement AbstractPager
	1. M_GetCurrentPageDataWithTB(context.Context, set.AbstractSet) error
	2. m_GetCurrentPageDataImpl(context.Context, set.AbstractSet) error
	3. M_SetPageSize(pageSize int32)
	4. M_SetRate(r float64)
	5. M_SetPageNumber(pageNumber int32)
	6. M_PageSize() int32
	7. M_PageNumber() int32
	8. M_HasNextThenToNext(set.AbstractSet) bool //是否有下一页？如果有就跳到下一页
	**/

	*pager.BasePager
	m_ecsHandler *EcsHandler
	m_req        *sdk_ecs.DescribeInstancesRequest
	// 后期考虑加入logger
}

func (ep *EcsPager) m_SetRequest() *sdk_ecs.DescribeInstancesRequest {
	//可以插一个log，看看当前是请求第几页的数据
	ep.m_req.PageNumber = sdk_req.NewInteger(int(ep.M_PageNumber()))
	ep.m_req.PageSize = sdk_req.NewInteger(int(ep.M_PageSize()))
	return ep.m_req
}

func (ep *EcsPager) m_GetCurrentPageDataImpl(ctx context.Context, s set.AbstractSet) error {
	resp, err := ep.m_ecsHandler.M_GetEcsHostGroupFromAlicloud(ep.m_SetRequest())
	//根据当前的PageNumber和PageSize，设置一个访问阿里云的request。然后拿着这个request去访问阿里云
	if err != nil {
		return err
	}
	s.M_Add(resp.M_TransferToTypeAny()...) //转换成any类型
	//s = resp
	//如果函数的最后一个参数是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数
	//如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数调用变参函数。
	ep.M_HasNextThenToNext(s)
	return nil
}

func CreateEcsPager(ecsHandler *EcsHandler) pager.AbstractPager {
	req := sdk_ecs.CreateDescribeInstancesRequest()
	ecsPager := &EcsPager{
		BasePager:    pager.CreateBasePager(),
		m_ecsHandler: ecsHandler,
		m_req:        req,
	}
	ecsPager.M_RegistFuncGetCurrentPageData(ecsPager.m_GetCurrentPageDataImpl)

	return ecsPager
}
