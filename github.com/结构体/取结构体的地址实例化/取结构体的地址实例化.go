package main

import "fmt"

//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("取结构体的地址的实例化")

	p := Person{}
	fmt.Printf("P === %v\n", p)  //P === { 0}
	fmt.Printf("P === %#v\n", p) //P === main.Person{name:"", age:0}
	p.name = "薛之谦"
	p.age = 36
	fmt.Printf("p === %v\n", p)  // p === {薛之谦 36}
	fmt.Printf("p === %#v\n", p) //p === main.Person{name:"薛之谦", age:36}
}
