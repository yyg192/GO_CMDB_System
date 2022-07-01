package main

import "fmt"

type AbstractTemplate interface {
	M_GetData() (int, error)
	m_GetData() (int, error)
}

type BaseTemplate struct {
	m_GetData func() (int, error)
}

func (bt *BaseTemplate) M_GetData() (int, error) {
	fmt.Println("预定义操作")
	return bt.m_GetData()
}

type Realizer struct {
	*BaseTemplate
}

func (r *Realizer) m_GetData() (int, error) {
	return 32, nil
}

func CreateRealizer() *Realizer {
	r := &Realizer{
		BaseTemplate: &BaseTemplate{},
	}
	r.BaseTemplate.m_GetData = r.m_GetData
	return r
}

func main() {
	r := CreateRealizer()
	ans, err := r.M_GetData()
	fmt.Println(ans)
	fmt.Println(err)
}
