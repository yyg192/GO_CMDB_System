package conf

const (
	P_testRegionId     string = "cn-hangzhou"
	P_testAccessKey    string = "LTAI5tQFjrUGN8KsZmCRrzfP"
	P_testAccessSecret string = "nSHMjM7fWYubDc4Awg5MzjCPTrZaq6"
)

func RegionId() string {
	return P_testRegionId
}

func AccessKey() string {
	return P_testAccessKey
}

func AccessSecret() string {
	return P_testAccessSecret
}
