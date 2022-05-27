package host

import (
	"context"

	"github.com/go-playground/validator"
)

/**
pkg/host/interface.go  提供统一服务接口:
	SaveHost QueryHost DescribeHost DeleteHost UpdateHost
服务接口需要的参数和返回值接口也在此处定义。这些接收的参数和返回的内容
在我们的后端业务中都是固定的，因为前端已经固定会收发这些内容了。
后面你可以去继承Service方法去实现自己的逻辑CRUD逻辑
***/

var (
	validate = validator.New()
)

type Service interface {
	SaveHost(context.Context, *Host) (*Host, error)
	//存储主机实例的接口
	//这里第一个参数为什么要传context.Context，goroutine里面也有一个context
	//可以用来取消正在运行的协程，这里的作用也是同理，假如我在调用这个Service服务方法
	//假如我中途需要取消，那就通过这个context.Context取消，把整个服务请求的链给取消

	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	//查询接口，查询请求通过QueryHostRequest抽象封装，便于后续我们可以更改查询请求的代码逻辑
	//后续如果想要变更查询请求的逻辑，就只需要改对应的代码不需要变更函数签名

	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	//为什么要定义查询 详细信息的接口，不直接使用QueryHost，还要专门做一个接口去查询
	//主机的详情？为了性能，为了缓存 01:30重新听

	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
	//删除一个主机实例，为什么要返回一个Host，用于审计或者全局事件广播之类的，我删除了一个实例
	//需要记录你删除的东西，或者把你删除的东西告诉其他方

	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
}

type QueryHostRequest struct {
	PageSize   uint64 `json:"page_size,omitempty"`   //请求的page的大小，比如前端一页20条数据
	PageNumber uint64 `json:"page_number,omitempty"` //当前查询的是第几页的host
	Keywords   string `json:"keywords"`
	//omitempty作用是在json数据结构转换时，当该字段的值为该字段类型的零值时，忽略该字段。
}

func NewDescribeHostRequestWithID(id string) *DescribeHostRequest {
	//这里留一个接口，万一以后DescribeHostRequest增加了新东西，直接在这里改
	return &DescribeHostRequest{
		Id: id,
	}
}

type DescribeHostRequest struct {
	Id string `json:"id"`
	//通过主机id查主机详情，不过也可以通过用户自定义的名字，邮箱之类的
}

type DeleteHostRequest struct {
	Id string `json:"id" validate:"required"`
}

type UpdateMode int

const (
	PUT   UpdateMode = iota //PUT是整个实例的更新
	PATCH                   //只改某个实例中的某个字段，比如name
)

type UpdateHostRequest struct {
	Id             string          `json:"id" validate:"required"`
	UpdateMode     UpdateMode      `json:"update_mode"`
	UpdateHostData *UpdateHostData `json:"data" validate:"required"`
}

func (req *UpdateHostRequest) Validate() error {
	return validate.Struct(req)
}

type UpdateHostData struct {
	// Update请求是不允许更新Base信息的，只能更新Resource和Describe
	*Resource
	*Describe
}

func (req *QueryHostRequest) Offset() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}
