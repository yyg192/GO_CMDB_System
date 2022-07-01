package host

func CreateHost() *Host {
	return &Host{
		BasicInformation:    &BasicInformation{},
		ResourceInformation: &ResourceInformation{},
		DescribeInformation: &DescribeInformation{},
	}
}

func CreateHostSet() *HostSet {
	return &HostSet{
		M_items: []*Host{}, //一定要初始化切片啊，不然报错!
	}
}

func (hs *HostSet) Add(items ...any) {
	for i := range items {
		hs.M_items = append(hs.M_items, items[i].(*Host))
	}
}

func (hs *HostSet) Length() int32 {
	return hs.M_total
}

func (hs *HostSet) M_TransferToTypeAny() []any {
	items := make([]any, hs.M_total)
	for _, item := range hs.M_items {
		items = append(items, item)
	}
	return items
}
