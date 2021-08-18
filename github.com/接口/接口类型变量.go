package main

import "fmt"

type Sayer interface {
	say()
}

type cat struct {
}

func (a cat) say() {

}

type dog struct {
}

func (d dog) say() {

}
func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪

	fmt.Println("接口类型变量")
}

/*
Tips： 观察下面的代码，体味此处_的妙用

// 摘自gin框架routergroup.go
type IRouter interface{ ... }

type RouterGroup struct { ... }

var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter
*/
