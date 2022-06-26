package ecs_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yyg192/GO_CMDB_System/provider/alicloud/connection"

	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/yyg192/GO_CMDB_System/provider/alicloud/ecs"
)

func Test_GetInstancesDescription(t *testing.T) {
	/**
	2022-6-26终于测试完毕
	**/
	Convey("测试函数: GetInstancesDescription", t, func() {
		Convey("获取cloud_client", func() {
			cloud_client := connection.CreateAliCloudClient(connection.Test_region_id, connection.Test_access_key, connection.Test_access_secret)
			cloud_client.EcsClientConnection()
			So(cloud_client, ShouldNotBeNil)
			Convey("获取ecs_handler", func() {
				ecs_handler := ecs.CreateEcsHandler(cloud_client.M_ecs_client)
				So(ecs_handler, ShouldNotBeNil)
				Convey("获取并打印来自云商的主机实例描述信息", func() {
					req := sdk_ecs.CreateDescribeInstancesRequest()
					host_set, err := ecs_handler.GetInstancesDescription(req)
					So(err, ShouldBeNil)
					fmt.Println("\n该用户拥有的Ecs实例数量为:", host_set.Total)

				})
			})
		})

	})
}
