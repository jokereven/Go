package main

import (
	"fmt"
	"reflect"
)

/*

	reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。
	reflect.Value与原始值之间可以互相转换。

	reflect.Value类型提供的获取原始值的方法如下：

	方法	说明
	Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
	Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
	Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
	Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
	Bool() bool	将值以 bool 类型返回
	Bytes() []bytes	将值以字节数组 []bytes 类型返回
	String() string	将值以字符串类型返回
	通过反射获取值

*/

func Reflectvalueof(x interface{}) {
	v := reflect.ValueOf(x)
	// fmt.Printf("%v\n %T\n", v, v) //鞠婧祎    reflect.Value
	//如何得到一个传入时候类型的变量
	k := v.Kind() //拿到值对应的类型类型
	switch k {
	case reflect.Float32:
		//把取到的值转换成一个int32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.String:
		ret := string(v.String())
		fmt.Printf("%v %T\n", ret, ret)
	}
}

func main() {
	//valueof
	var str = "鞠婧祎"
	Reflectvalueof(str)
}
