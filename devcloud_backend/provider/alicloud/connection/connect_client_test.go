package connection_test

import (
	"testing"

	. "github.com/yyg192/GO_CMDB_System/provider/alicloud/connection"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_EcsClientConnection(t *testing.T) {

	Convey("测试函数: EcsClientConnection", t, func() {
		cloud_client := CreateAliCloudClient(
			Test_region_id,
			Test_access_key,
			Test_access_secret,
		)
		client, err := cloud_client.EcsClientConnection()
		Convey("获取客户Ecs服务连接的句柄", func() {
			So(client, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		cloud_client.PrintAliCloudClientInfo()
	})
}
