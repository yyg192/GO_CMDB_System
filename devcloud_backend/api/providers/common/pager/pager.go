package pager

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/yyg192/GO_CMDB_System/api/providers/common/set"
)

type AbstractPager interface {
	M_GetCurrentPageDataWithTB(context.Context, set.AbstractSet) error
	m_GetCurrentPageDataImpl(context.Context, set.AbstractSet) error
	M_SetPageSize(pageSize int32)
	M_SetRate(r float64)
	M_SetPageNumber(pageNumber int32)
	M_PageSize() int32
	M_PageNumber() int32
	M_HasNextThenToNext(set.AbstractSet) bool //是否有下一页？如果有就跳到下一页
}

type GetCurrentPageDataSign = func(context.Context, set.AbstractSet) error
type BasePager struct { // 继承自AbstractPager
	m_pageSize   int32
	m_pageNumber int32
	//m_hasNext    bool
	tb *tokenbucket.Bucket
	//需要继承者自己实现的函数
	GetCurrentPageDataFuncptr GetCurrentPageDataSign
}

func CreateBasePager() *BasePager {
	return &BasePager{
		m_pageSize:   20,
		m_pageNumber: 1,
		//m_hasNext:    true,
		tb: tokenbucket.NewBucketWithRate(1, 1),
	}
}

// func (bp *BasePager) M_HasNext() bool {
// 	bp.tb.Wait(1) //调用GetCurrentPageData的时候是需要
// 	return bp.m_hasNext
// }
// func (bp *BasePager) m_GetCurrentPageDataImpl(ctx context.Context, s set.AbstractSet) error {
// 	//带有令牌桶流量控制的 GetCurrentPageData
// 	if bp.m_GetCurrentPageDataImplPtr == nil {
// 		return fmt.Errorf("你必须要实现GetCurrentPageData方法并调用M_PassFuncGetCurrentPageData(GetCurrentPageData)")
// 	}
// 	return bp.m_GetCurrentPageDataImplPtr(ctx, s)

// }

func (bp *BasePager) M_SetPageSize(pageSize int32) {
	bp.m_pageSize = pageSize
}

func (bp *BasePager) M_SetPageNumber(pageNumber int32) {
	bp.m_pageNumber = pageNumber
}

func (bp *BasePager) M_SetRate(rate float64) {
	bp.tb.SetRate(rate)
}

func (bp *BasePager) M_PageSize() int32 {
	return bp.m_pageSize
}

func (bp *BasePager) M_PageNumber() int32 {
	return bp.m_pageNumber
}

func (bp *BasePager) M_GetCurrentPageDataWithTB(ctx context.Context, s set.AbstractSet) error {
	err := bp.m_GetCurrentPageDataImpl(ctx, s)
	return err
}

func (bp *BasePager) m_GetCurrentPageDataImpl(ctx context.Context, s set.AbstractSet) error {
	if bp.GetCurrentPageDataFuncptr == nil {
		return fmt.Errorf("you need to call M_RegistFuncGetCurrentPageData to regist the main body of GetCurrentPageData")
	}
	return bp.GetCurrentPageDataFuncptr(ctx, s)
}

func (bp *BasePager) M_RegistFuncGetCurrentPageData(funcPtr GetCurrentPageDataSign) {
	bp.GetCurrentPageDataFuncptr = funcPtr
}

func (bp *BasePager) M_HasNextThenToNext(set set.AbstractSet) bool {
	// Not sure?
	if int32(set.M_Length()) < bp.m_pageSize {
		return false
	} else {
		bp.m_pageNumber++
		return true
	}
}
