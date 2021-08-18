package main

import "fmt"

//使用值接收者与使用指针接受者的区别

type Mover interface {
	move()
}

type Person struct {
	Name string
	age  int
}

/*
	Person实现move()方法
	使用值接收者实现接口
*/

// func (p Person) move() {
// 	fmt.Printf("%s在跑\n", p.Name)
// }

//使用指针接受者实现接口 当使用指针接收者实现接口时不可以使用person类型的值
//只有接口类型的值为指针才可以使用
func (p *Person) move() {
	fmt.Printf("%s再跑", p.Name)
}

func main() {
	// var m1 Mover
	var m2 Mover
	// p1 := Person{
	// 	Name: "鞠婧祎",
	// 	age:  32,
	// }
	p2 := &Person{ //p2为Person类型的指针
		Name: "杨超越",
		age:  28,
	}
	// m1 = p1 //无法赋值因为是person类型的值
	m2 = p2
	// m1.move()
	m2.move()
	// fmt.Println(m1)
	fmt.Println(m2)
}
