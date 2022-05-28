package http

import (
	"CMDBProject/api/pkg"
	"CMDBProject/api/pkg/host"
	"fmt"
	"net/http"
	"strconv"

	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

var (
	api = &handler{}
)

const ( //设置分页的默认值
	defaultPageSize   = 20
	defaultPageNumber = 1
)

type handler struct {
	service host.Service //不能声明一个接口的指针
	//log     *logrus.Logger
}

func (h *handler) Config() error {
	h.service = pkg.Host
	if pkg.Host == nil {
		return fmt.Errorf("dependence service host not ready")
	}
	//h.log =  ????
	return nil
}

func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query() //获取他的query 字符串
	//https://baijiahao.baidu.com/s?id=1603848351636567407&wfr=spider&for=pc
	// ?之后和#之前的都是Query字符串
	var (
		pageSize   = defaultPageSize
		pageNumber = defaultPageNumber
	)
	pageSizeStr := qs.Get("page_size")
	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}
	pageNumberStr := qs.Get("page_number")
	if pageNumberStr != "" { //一般不会出现这种参数为空的
		pageNumber, _ = strconv.Atoi(pageNumberStr)
	}
	req := &host.QueryHostRequest{
		PageSize:   uint64(pageSize),
		PageNumber: uint64(pageNumber),
	}
	set, err := h.service.QueryHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ins := host.NewDefaultHost()
	if err := GetDataFromRequest(r, ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins, err := h.service.SaveHost(r.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) DescribeHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewDescribeHostRequestWithID(ps.ByName("id"))
	//
	set, err := h.service.DescribeHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewDeleteHostRequestWithID(ps.ByName("id"))
	set, err := h.service.DeleteHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PutHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewUpdateHostRequest(ps.ByName("id"))
	//http://127.0.0.1:6050/hosts/c4mb07u8123456
	//id就是c4mb07u8123456
	if err := GetDataFromRequest(r, req.UpdateHostData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) PatchHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewUpdateHostRequest(ps.ByName("id"))
	req.UpdateMode = host.PATCH
	if err := GetDataFromRequest(r, req.UpdateHostData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.GET("/hosts", api.QueryHost)
	r.POST("/hosts", api.CreateHost)
	r.GET("/hosts/:id", api.DescribeHost)
	r.DELETE("/hosts/:id", api.DeleteHost)
	r.PUT("/hosts/:id", api.PutHost)
	r.PATCH("/hosts/:id", api.PatchHost)
}
