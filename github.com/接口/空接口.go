package main

import "fmt"

/*
	空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

	空接口类型的变量可以存储任意类型的变量。
*/

//空接口一帮不提前定义
type null interface{}

func main() {
	var null interface{} //定义一个空接口类型
	null = false
	fmt.Println(null)
	null = "鞠婧祎"
	fmt.Println(null)
}
