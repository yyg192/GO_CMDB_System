package set

type AbstractSet interface {
	M_Add(...any)               //往里面添加任意数据类型的元素
	M_Length() int32            //返回元素个数
	M_TransferToTypeAny() []any //将集合内所有元素读转换为any类型
}
