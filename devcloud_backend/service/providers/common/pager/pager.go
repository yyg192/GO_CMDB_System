package pager

import "context"

type AbstractSet interface {
	Add(...interface{})
	Length() int64
}

type AbstractPager interface {
	M_HasNext() bool
	M_GetCurrentPageData(context.Context, AbstractSet) //只要实现了Add和Length方法的结构体就有资格作为他的参数
}

type BasePager struct {
	M_pageSize int32
}
