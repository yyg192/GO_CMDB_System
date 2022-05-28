package pkg

/***
pkg包用来管理所有的实例，这个文件相当于一个IOC容器，
***/
import (
	"CMDBProject/api/pkg/host"
)

var (
	Host host.Service
)
