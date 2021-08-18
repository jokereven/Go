package main

import (
	"fmt"
	"reflect"
)

/*
	在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。 举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。
*/

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type_of_v:%v\n ", v)
	fmt.Printf("%T\n", v) // *reflect.rtype
	fmt.Println(v, v.Name(), v.Kind())
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

/*
	Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。

在reflect包中定义的Kind类型如下：

type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
*/
