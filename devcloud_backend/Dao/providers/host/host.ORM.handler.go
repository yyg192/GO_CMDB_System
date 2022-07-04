package host

func CreateHost() *Host {
	return &Host{
		BasicInformation:    &BasicInformation{},
		ResourceInformation: &ResourceInformation{},
		DescribeInformation: &DescribeInformation{},
	}
}

func CreateHostSet() *HostSet {
	//继承自AbstractSet
	return &HostSet{
		M_items: []*Host{}, //一定要初始化切片啊，不然报错!
		M_total: 0,
	}
}

func (hs *HostSet) M_Add(items ...any) {
	for i := range items {
		hs.M_items = append(hs.M_items, items[i].(*Host))
		//hs.M_items是 []*Host 类型
	}
	/**
	for _, item := range items {
		hs.M_items = append(hs.M_items, items[i].(*Host)) 这样会直接panic！
		//这个bug找的好辛苦啊！回头研究一下为什么？？？？？
	}
	**/
}

func (hs *HostSet) M_Length() int32 {
	return hs.M_total
}

func (hs *HostSet) M_TransferToTypeAny() (items []any) {
	//items := make([]any, hs.M_total)
	for i := range hs.M_items {
		items = append(items, hs.M_items[i])
	}
	return
}

/**
为什么这个函数就会报错，上面的不会啊 ！！！！
func (hs *HostSet) M_TransferToTypeAny() []any {
	items := make([]any, hs.M_total)
	for i := range hs.M_items {
		items = append(items, hs.M_items[i])
	}
	return items
}
**/
