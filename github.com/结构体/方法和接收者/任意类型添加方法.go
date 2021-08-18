package main

import "fmt"

/*
在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
*/

//MyInt 将int定义为自定义MyInt类型
type MyInt int

func (m MyInt) SayHello() {
	fmt.Println("任意类型的添加方法")
}
func main() {
	var m1 MyInt
	m1.SayHello() //任意类型的添加方法
	m1 = 996
	m1.SayHello()                                   //任意类型的添加方法
	fmt.Printf("m1-value:%#v m1=type:%T\n", m1, m1) //m1-value:996 m1=type:main.MyInt
}
