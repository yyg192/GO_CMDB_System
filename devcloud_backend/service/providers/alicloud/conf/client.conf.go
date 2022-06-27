package conf

const (
	P_testRegionId     string = "cn-hangzhou"
	P_testAccessKey    string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	P_testAccessSecret string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
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
