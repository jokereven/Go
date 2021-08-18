package main

import "fmt"

/*
定义函数类型
我们可以使用type关键字来定义一个函数类型，具体格式如下：

type calculation func(int, int) int
上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。
*/

/*
add和sub都能赋值给calculation类型的变量。
var c calculation
c = add
*/

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	// var a1 = add(1, 2) //3
	// var a2 = sub(1, 2) //-1
	// fmt.Println(a1, a2)

	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}
