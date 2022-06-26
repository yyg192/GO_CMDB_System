package host

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{}, //一定要初始化切片啊，不然报错！
	}
}

func NewDefaultHost() *Host {
	return &Host{
		Base:     &Base{},
		Resource: &Resource{},
		Describe: &Describe{},
	}
}
