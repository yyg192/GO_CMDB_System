package resource

import "github.com/yyg192/GO_CMDB_System/api/providers/set"

func CreateResource() *Resource {
	return &Resource{
		BasicInformation:  &BasicInformation{},
		DetailInformation: &DetailInformation{},
	}
}

func CreateResourceSet() set.AbstractSet {
	return &ResourceSet{
		M_items: []*Resource{}, //一定要初始化切片啊，不然报错!
		M_total: 0,
	}
}

func (rs *ResourceSet) M_Add(items ...any) {
	for i := range items {
		rs.M_items = append(rs.M_items, items[i].(*Resource))
		//hs.M_items是 []*Host 类型
	}
}

func (rs *ResourceSet) M_Length() int32 {
	return rs.M_total
}

func (rs *ResourceSet) M_TransferToTypeAny() (items []any) {
	//items := make([]any, hs.M_total)
	for i := range rs.M_items {
		items = append(items, rs.M_items[i])
	}
	return
}
