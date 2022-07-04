package connection

import (
	"testing"

	//. "github.com/yyg192/GO_CMDB_System/BO/provider/clouds/alicloud/connection"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yyg192/GO_CMDB_System/service/providers/alicloud/conf"
)

func Test_EcsClientConnection(t *testing.T) {
	Convey("测试函数: EcsClientConnection", t, func() {
		cloud_client := CreateAliCloudClient(
			conf.RegionId(),
			conf.AccessKey(),
			conf.AccessSecret(),
		)
		client, err := cloud_client.M_EcsClientConnection()
		Convey("获取客户Ecs服务连接的句柄", func() {
			So(client, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		PrintAliCloudClientInfo(cloud_client)
	})
}
