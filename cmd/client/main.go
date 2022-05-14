package main

import "github.com/yyg192/Go_FileStation/cmd/client/cmd"

//导入cmd包，cmd包下面的全局变量也会被导入进来，cmd包下面的init函数也会被执行。

func main() {
	cmd.Execute()
}
