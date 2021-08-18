package main

import "fmt"

/*
高阶函数分为函数作为参数和函数作为返回值两部分。
*/

/*
函数作为参数
*/

func add(x, y int) int {
	return x + y
}

func num(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	fmt.Println("函数作为参数")
	ret := num(996, 985, add)
	fmt.Println(ret) //30
}
