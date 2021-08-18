package main

import "fmt"

/*
闭包
闭包指的是一个函数和与其相关的引用环境组合而成的实体。
简单来说，闭包=函数+引用环境。 首先我们来看一个例子：
*/

//定义一个函数返回为一个匿名函数
func val(name string) func() {
	// name := "鞠婧祎"	//直接使用或到val(name string)
	return func() {
		fmt.Println("闭包-01", name)
	}
}

func main() {
	fmt.Println("闭包")

	r := val("鞠婧祎")
	r() //执行函数内部的匿名函数
}
