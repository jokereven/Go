package main

import "fmt"

/*
	函数
Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。

函数定义
Go语言中定义函数使用func关键字，具体格式如下：

func 函数名(参数)(返回值){
    函数体
}
其中：

函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
函数体：实现指定功能的代码块。
*/

func Hello() {
	fmt.Println("Hello World")
}

func Name(Name string) {
	fmt.Println("我的名字叫" + Name)
}

func Sum(a int, b int) int {
	sum := a + b
	return sum
}

//Go语言参数类型简写
func SUM(a, b int) int {
	sum := a + b
	return sum
}

func retSum(a int, b int) (ret int) {
	ret = a + b
	return
}

func main() {
	Hello()

	name := "鞠婧祎"
	Name(name)

	Name("鞠婧祎")

	index := Sum(1314, 996)
	fmt.Println(index)

	retsum := retSum(520, 1314)
	fmt.Println(retsum)
}
