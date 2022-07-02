## 2022-06-06
原来的一把梭架构太粗鄙了，随着新功能的不断增加，代码量的增加，原来的项目组织架构感觉混乱不堪，无法继续推进，还是重新把后端设计一下吧。

## 2022-06-14
稍微开发了一下前端，基本界面样式算是出来了，当然还有以下事情还没有完成
Todolist:
1. 登录界面的背景还没有搞好。
2. Layout布局可能还有点问题，尤其是侧边栏导航的收缩过程中的内容展示。
3. 还有一些美工计划，比如CMDB界面是春天的背景，Workflow是秋天的背景，按照春夏秋冬分别设计一下，感觉极具美感，嘿嘿。

## 2022-06-26
后端重构了一小下，采用了MVC组织架构，目前完成了调用阿里云的sdk获取Ecs实例描述信息的功能
Todolist:
1. 一次调用返回所有的Ecs实例描述信息，但是我希望每次获取能只获取一页！不要全都拿回来
2. 接下来把数据库增删改查逻辑给整一整

## 2022-06-27
后端又重构了一下，最终确定了组织架构，目前我结合阿里巴巴编程规约制定了一个自己的Go编程规约，暂时试验一下看好不好用。

## 2022-07-02
太晚了先睡了，还有个bug没改完，bug出现在ecspager.go文件中的第50行 s.M_Add(resp.M_TransferToTypeAny()...)
我这几天基本上又重新的组织了一下项目组织架构，更加规范化。接下来是要完成ecs_test.go的测试内容，利用pgaer获取云商数据。另外有几个todo事项。
1. 好好看看Test_GetFIrstPageDataFromAliCloud这个函数，尤其是每次创建完cloud_client后还要执行M_EcsClientConnection()，要是忘记的话就会报错，而且这极度依赖用户操作。这是不行的！！！要把这个连接云客户端的操作想办法给塞进CreateAliCloudClient函数里面，
2.  另外我感觉最终的目的其实就是为了获取ecsPager，我觉得可以完全的从创建cloud_client，连接cloudClient，创建ecsHandler以及ecsPager这几个步骤合在一起，就不用每次都写这么多内容了，怪麻烦的。
3.  另外回头把我自己探索出来的模板模式给写进我的go编程规约里面。

## 2022-07-02
Test_GetFirstPageDataFromAliCloud这个测试函数终于过了，到目前为止代码预期功能都正常实现了，虽然bug解决了，但是也是有点误打误撞解决的，比如下面这些问题我都还没搞懂。
``` go
func (hs *HostSet) M_TransferToTypeAny() (items []any) {
	//items := make([]any, hs.M_total)
	for i := range hs.M_items {
		items = append(items, hs.M_items[i])
	}
	return
}

/**
为什么这个函数就会报错，上面的不会啊 ！！！！
func (hs *HostSet) M_TransferToTypeAny() []any {
	items := make([]any, hs.M_total)
	for i := range hs.M_items {
		items = append(items, hs.M_items[i])
	}
	return items
}
**/
```
还有下面这个奇怪的bug，虽然改对了，但是仍然不知道错误版本有什么问题，回头再研究一下！
``` go
func (hs *HostSet) M_Add(items ...any) {
	for i := range items {
		hs.M_items = append(hs.M_items, items[i].(*Host))
		//hs.M_items是 []*Host 类型
	}
	/**
	for _, item := range items {
		hs.M_items = append(hs.M_items, items[i].(*Host)) 这样会直接panic！
		//这个bug找的好辛苦啊！回头研究一下为什么？？？？？
	}
	**/
}
```
然后几个todo list没完成
1.  另外我感觉最终的目的其实就是为了获取ecsPager，我觉得可以完全的从创建cloud_client，连接cloudClient，创建ecsHandler以及ecsPager这几个步骤合在一起，就不用每次都写这么多内容了，怪麻烦的。
2.  另外回头把我自己探索出来的模板模式给写进我的go编程规约里面。
