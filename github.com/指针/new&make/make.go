package main

import "fmt"

/*
make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：

func make(t Type, size ...IntegerType) Type
make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。这个我们在上一章中都有说明，关于channel我们会在后续的章节详细说明。

本节开始的示例中var b map[string]int只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值：
*/

func main() {
	var a map[string]int
	a = make(map[string]int, 4)
	a["鞠婧祎"] = 24
	fmt.Println(a) //map[鞠婧祎:24]
}

/*
ew与make的区别
二者都是用来做内存分配的。
make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
