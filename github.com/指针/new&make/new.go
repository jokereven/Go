package main

import "fmt"

/*
new是一个内置的函数，它的函数签名如下：

func new(Type) *Type
其中，

Type表示类型，new函数只接受一个参数，这个参数是一个类型
*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：
*/

func main() {
	// a := new(int)
	// b := new(bool)
	// fmt.Printf("a:%T\n", a) //a:*int
	// fmt.Printf("b:%T\n", b) //b:*bool
	// fmt.Println(a)          //0xc0000140a8
	// fmt.Println(b)          //0xc0000140c0
	// fmt.Println(*a)         //0
	// fmt.Println(*b)         //false
	/*
		本节开始的示例代码中var a *int只是声明了一个指针变量a但是没有初始化，
		指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
		应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：
	*/
	var f *int
	fmt.Println(f) //nil空的指针
	f = new(int)   //初始化
	fmt.Println(f) //0xc0000ac090
	*f = 996
	fmt.Printf("f:%T\n", *f) //f:int
	fmt.Println(*f)          //996
}
