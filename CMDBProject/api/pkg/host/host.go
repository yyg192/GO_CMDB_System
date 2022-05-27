package host

import (
	"time"

	"github.com/imdario/mergo"
)

type Vendor int

/***
和云主机有关的所有信息
***/

//为什么这里要费劲写一个Vendor，因为其他人调用我的api，发现是个int，对方不知道
//我这个int应该传多少，范围是什么，代表什么含义，如果我这里type Vendor int
//对方就知道我这个int是用在哪个地方，每个值对应什么含义

const (
	PrivateIDC Vendor = iota
	Tencent
	Aliyun
	Huawei
)

type Host struct {
	//一个Host由下面这三段信息组成
	*Base
	*Resource
	*Describe
}

func (h *Host) Update(res *Resource, desc *Describe) {
	h.Resource.UpdateAt = time.Now().UnixNano() / 1000000
	h.Resource = res
	h.Describe = desc
}

func (h *Host) Patch(res *Resource, desc *Describe) error {
	h.Resource.UpdateAt = time.Now().UnixNano() / 1000000
	if res != nil {
		err := mergo.MergeWithOverwrite(h.Resource, res)
		if err != nil {
			return err
		}
	}
	if desc != nil {
		err := mergo.MergeWithOverwrite(h.Describe, desc)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func NewDefaultHost() *Host {
	return &Host{
		Base:     &Base{},
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

func NewDefaultHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{}, //防止Items空指针，还是初始化一个空值吧
	}
}

func (hs *HostSet) Add(h *Host) {
	hs.Items = append(hs.Items, h)
}

/**
Base
云资源的通用基础信息，方便后续检索
**/
type Base struct {
	Id           string `json:"id"`            //全局唯一id
	SyncAt       int64  `json:"sync_at"`       //同步时间
	Vendor       Vendor `json:"vendor"`        //厂商
	Region       string `json:"region"`        //地域
	Zone         string `json:"zone"`          //区域
	CreateAt     int64  `json:"CreateAt"`      //创建时间
	InstanceId   string `json:"instance_id"`   //实例ID
	ResourceHash string `json:"resource_hash"` //基础数据Hash 相当于把resource字段做一个哈希，用来判断他有没有发生变化
	DescribeHash string `json:"describe_hash"` //描述数据Hash
}

/**
跟云实例有关的数据
**/
type Resource struct {
	ExpireAt    int64             `json:"expire_at"`    //过期时间
	Category    string            `json:"category"`     //种类
	Type        string            `json:"type"`         //规格
	Name        string            `json:"name"`         //名称
	Description string            `json:"description"`  //描述
	Status      string            `json:"status"`       //服务商中的状态 running or ?
	Tags        map[string]string `json:"tags"`         //标签
	UpdateAt    int64             `json:"update_at"`    //更新时间
	SyncAccount string            `json:"sync_account"` //同步的帐号 一个公司下面可能有很多帐号，需要进行区分，要知道这个资源是哪个帐号同步过来的
	PublicIP    string            `json:"public_ip"`    //公网IP
	PrivateIP   string            `json:"private_ip"`   //内网IP
	PayType     string            `json:"pay_type"`     //实例付费方式
}

//不通用的信息，不同的资源资产有不一样的信息，比如一些资产没有GPU的信息。
type Describe struct {
	ResourceId              string `json:"resource_id"`                // 关联Resource
	CPU                     int    `json:"cpu"`                        // 核数
	Memory                  int    `json:"memory"`                     // 内存
	GPUAmount               int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec                 string `json:"gpu_spec"`                   // GPU类型
	OSType                  string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName                  string `json:"os_name"`                    // 操作系统名称
	SerialNumber            string `json:"serial_number"`              // 序列号
	ImageID                 string `json:"image_id"`                   // 镜像ID
	InternetMaxBandwidthOut int    `json:"internet_max_bandwidth_out"` // 公网出带宽最大值，单位为 Mbps
	InternetMaxBandwidthIn  int    `json:"internet_max_bandwidth_in"`  // 公网入带宽最大值，单位为 Mbps
	KeyPairName             string `json:"key_pair_name,omitempty"`    // 秘钥对名称
	SecurityGroups          string `json:"security_groups"`            // 安全组  采用逗号分隔
}
