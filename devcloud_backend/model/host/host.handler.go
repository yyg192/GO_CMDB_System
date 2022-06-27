package host

func (ag *AccountGetter) GetAccountId() string {
	return ag.accountId
}

func (hs *HostSet) Add(items ...any) {
	for i := range items {
		hs.Items = append(hs.Items, items[i].(*Host))
	}
}

func (hs *HostSet) Length() int32 {
	return hs.Total
}
