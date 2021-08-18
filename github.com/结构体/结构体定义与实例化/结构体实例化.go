package main

import "fmt"

/*
只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

var 结构体实例 结构体类型
基本实例化
*/

type Persion struct {
	name, dream string
	age         int
}

func main() {
	fmt.Println("结构体的实例化")
	var joker_xue Persion
	joker_xue.name = "薛之谦"
	joker_xue.dream = "世界和平"
	joker_xue.age = 36
	fmt.Printf("joker_xue === %v\n", joker_xue)  //joker_xue === {薛之谦 世界和平 36}
	fmt.Printf("joker_xue === %#v\n", joker_xue) //joker_xue === main.Persion{name:"薛之谦", dream:"世界和平", age:36}
}

/*
我们通过.来访问结构体的字段（成员变量）,例如p1.name和p1.age等。
*/
