package main

import "fmt"

//一个类型实现多个接口

//接口的嵌套

type Kiku interface{
	Mover
	Play
}
type Mover interface {
	move()
}

type Play interface{
	play()
}

type Person struct {
	Name string
	age  int
}

func (pmove Person) move() {
	fmt.Printf("%s再跑\n", pmove.Name)
}

func (pplay Person) play(){
	fmt.Printf("%s在跳舞\n",pplay.Name);
}
func main() {
	var m1 Mover
	var m2 Play
	p1 := Person{
		Name: "鞠婧祎",
		age:  32,
	}
	m1 = p1 //无法赋值因为是person类型的值
	m2 = p1
	m1.move()
	m2.play();
	fmt.Println(m1)
	fmt.Println(m2)

	var m3 Kiku
	m3 = p1
	m3.move()
	m3.play()
}
