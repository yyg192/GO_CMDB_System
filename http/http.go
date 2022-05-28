package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// BodyMaxContenxLength body最大大小 默认64M
	BodyMaxContenxLength int64 = 1 << 26
)

// GetDataFromRequest todo
func GetDataFromRequest(r *http.Request, v interface{}) error {
	body, err := ReadBody(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

// ReadBody 读取Body当中的数据，就是http报文最后的数据部分
func ReadBody(r *http.Request) ([]byte, error) {
	// 检测请求大小
	if r.ContentLength == 0 {
		return nil, fmt.Errorf("request body is empty")
	}
	if r.ContentLength > BodyMaxContenxLength { //传的数据太大了
		return nil, fmt.Errorf(
			"the body exceeding the maximum limit, max size %dM",
			BodyMaxContenxLength/1024/1024)
	}

	// 读取body数据
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("read request body error, %s", err))
	}

	return body, nil
}
