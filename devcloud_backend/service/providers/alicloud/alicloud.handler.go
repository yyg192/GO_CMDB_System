package alicloud

/**
能够操控阿里云的所有资源！
**/
import (
	"sync"

	"github.com/yyg192/GO_CMDB_System/service/providers/alicloud/conf"
	"github.com/yyg192/GO_CMDB_System/service/providers/alicloud/connection"
)

var (
	handler         *Handler
	loadHandlerOnce sync.Once
)

func H() *Handler { //用单例懒汉模式实现，第一次调用Handler的时候会初始化，后面再调用就不会了
	loadHandlerOnce.Do(
		func() {
			var err error
			handler, err = CreateHandler(conf.RegionId(),
				conf.AccessKey(),
				conf.AccessSecret())
			if err != nil {
				panic("handle initialization failed")
			}
		})
	return handler
}

type Handler struct {
	m_cloudClient *connection.AliCloudClient
	m_accountId   string //来自云商上面记录的账户Id
}

// conf.RegionId(), conf.AccessKey(), conf.AccessSecret()
func CreateHandler(regionId, accessKey, accessSecret string) (*Handler, error) {
	alicloudClient := connection.CreateAliCloudClient(regionId, accessKey, accessSecret)
	accountId, err := alicloudClient.M_account()
	if err != nil {
		return nil, err
	}
	return &Handler{
		m_cloudClient: alicloudClient,
		m_accountId:   accountId,
	}, nil
}

/**
func (o *Operator) HostOperator() provider.HostOperator {
	c, err := o.client.EcsClient()
	if err != nil {
		panic(err)
	}
	op := ecs.NewEcsOperator(c)
	op.WithAccountId(o.account)
	return op
}

func (o *Operator) BillOperator() provider.BillOperator {
	c, err := o.client.BssClient()
	if err != nil {
		panic(err)
	}
	return bss.NewBssOperator(c)
}

func (o *Operator) RdsOperator() provider.RdsOperator {
	c, err := o.client.RdsClient()
	if err != nil {
		panic(err)
	}
	return rds.NewRdsOperator(c)
}

**/
