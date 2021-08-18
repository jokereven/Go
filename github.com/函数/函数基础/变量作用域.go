package main

import "fmt"

/*
全局变量
全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。
在函数中可以访问到全局变量。
*/

/*
局部变量
局部变量又分为两种： 函数内定义的变量无法在该函数外使用，
例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：
*/

/*
如果局部变量和全局变量重名，优先访问局部变量。
*/

// var name = "鞠婧祎"

// func actionScope() {
// 	var val = "杨超越"
// 	fmt.Println(val)
// }

var num int = 1314

func Num() {
	var num = 520
	fmt.Println(num)
}

func main() {
	fmt.Println("作用域")
	// fmt.Println(val) //undefined

	Num()
}
