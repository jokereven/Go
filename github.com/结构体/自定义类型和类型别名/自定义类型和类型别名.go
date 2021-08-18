package main

import "fmt"

/*
Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
*/

/*
类型别名和自定义类型
自定义类型
在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型， Go语言中可以使用type关键字来定义自定义类型。

自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。
*/

/*
类型别名
类型别名是Go1.9版本添加的新功能。

类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
*/

type Myint int

// 类型别名
type Theint = int

func main() {
	var i Myint
	fmt.Printf("type of Intmy:%T\n", i) //type of Intmy:main.Myint
	var x Theint
	fmt.Printf("type of Theint:%T\n", x) //type of Theint:int
}

/*
通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
*/
