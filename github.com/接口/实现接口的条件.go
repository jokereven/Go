package main

import "fmt"

/*
一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。
*/

//定义一个Sayer的接口
type Sayer interface {
	say()
}

//定义猫🐱和狗的结构体
type Cat struct {
}
type Dog struct {
}

//Cat🐱实现Sayer接口
func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

//Dog实现Sayer接口
func (d Dog) Say() {
	fmt.Println("汪汪汪")
}
func main() {
	fmt.Println("实现接口")
}
