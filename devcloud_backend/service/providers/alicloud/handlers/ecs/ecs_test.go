package ecs_test

import (
	"fmt"
	"testing"

	sdk_ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yyg192/GO_CMDB_System/service/providers/alicloud/conf"
	"github.com/yyg192/GO_CMDB_System/service/providers/alicloud/connection"
	"github.com/yyg192/GO_CMDB_System/service/providers/alicloud/handlers/ecs"
)

func Test_GetAllInstancesDescription(t *testing.T) {
	/**
	2022-6-26终于测试完毕
	**/
	Convey("测试函数: GetInstancesDescription", t, func() {
		Convey("获取cloud_client", func() {
			cloud_client := connection.CreateAliCloudClient(conf.RegionId(), conf.AccessKey(), conf.AccessSecret())
			cloud_client.M_EcsClientConnection()
			So(cloud_client, ShouldNotBeNil)
			Convey("获取ecs_handler", func() {
				ecs_handler := ecs.CreateEcsHandler(cloud_client.M_ecs_client)
				So(ecs_handler, ShouldNotBeNil)
				Convey("获取并打印来自云商的主机实例描述信息", func() {
					req := sdk_ecs.CreateDescribeInstancesRequest()
					host_set, err := ecs_handler.M_GetAllInstancesDescriptionFromAlicloud(req)
					So(err, ShouldBeNil)
					fmt.Println("\n该用户拥有的Ecs实例数量为:", host_set.Total)
				})
			})
		})

	})
}
