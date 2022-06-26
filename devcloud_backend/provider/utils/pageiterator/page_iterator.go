package pageiterator

import (
	"context"

	sdk_tokenbucket "github.com/infraboard/mcube/flowcontrol/tokenbucket"
)

type Set interface {
	Append(...interface{})
	Length() int64
}

//分页迭代器的接口
type I_PageIterator interface {
	//
	HasNextPage() bool
	SetPageSize(page_size int64)
	SetRate(r float64)
	GetCurrentPageData(context.Context, Set) error
}

type PageIterator struct {
	page_size      int64
	page_number    int64
	hast_next_page bool
	tooken_bucket  *sdk_tokenbucket.Bucket
}

func (page_iter *PageIterator) HasNextPage() bool {
	page_iter.tooken_bucket.Wait(1) //访问之前就要先从桶里拿一个令牌出来，
	// Wait就是等一个令牌操作，有就拿没就等
	return true
}
