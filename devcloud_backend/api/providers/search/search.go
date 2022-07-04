package search

/**
我有一个全局资源检索的功能
**/

import (
	"context"

	"github.com/yyg192/GO_CMDB_System/api/providers/set"
)

type Type int

type Vendor int

const (
	VendorAliCloud Vendor = iota
	VendorTencentCloud
	VendorHuaWeiCloud
)

const (
	Unsuport Type = iota
	EcsResource
	RdsResource
)

type SearchRequest struct {
	Vendor       Vendor //int range: 0,1,2
	ResourceType Type   //int range: 0,1,2
}

type Search interface {
	M_Search(context.Context, *SearchRequest) (*set.AbstractSet, error)
}
