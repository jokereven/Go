package main

import (
	"fmt"
	"reflect"
)

/*
	在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的(我们在上一篇接口的博客中有介绍相关概念)。 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。
*/

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type_of_v:%v\n ", v)
	fmt.Printf("%T\n", v) // *reflect.rtype
}

type Cat struct{}
type Dog struct{}

func main() {
	fmt.Println("reflect包")
	var num int = 996
	var str string = "鞠婧祎"
	reflectType(num) //type_of_v:int
	reflectType(str) // type_of_v:string

	//结构体
	var c Cat
	reflectType(c) //type_of_v:main.Cat
	var d Dog
	reflectType(d) //type_of_v:main.Cat
}
