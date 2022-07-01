/**
这个功能怎么实现啊，感觉golang有的地方还挺鸡肋的，要是interface能实现方法就好了
**/

package nothing

import "fmt"

type AbstractClass interface {
	M_GetData() error
	m_GetData() error
}

func (ac *AbstractClass) M_GetData() {
	fmt.Println("预定义的操作")
	return m_GetData()
}

type BaseClass struct {
}

/* 由继承者决定怎么 */
func (bc *BaseClass) m_GetData() {
	fmt.Println("继承者实现GetData的具体逻辑")
}

func main() {
	bc := &BaseClass{}
	bc.M_GetData()
}
