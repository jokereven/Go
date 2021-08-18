package main

import "fmt"

//我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("指针类型结构体")
	var joker_xue = new(Person)
	joker_xue.name = "薛之谦"
	joker_xue.age = 36
	fmt.Printf("type of joker_xue === %#v\n", joker_xue) //joker_xue === {薛之谦 36}
	fmt.Printf("joker_xue === %#v\n", joker_xue)         //joker_xue === struct { name string; age int }{name:"薛之谦", age:36}
}

//需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。
