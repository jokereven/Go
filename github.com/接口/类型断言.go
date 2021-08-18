package main

import "fmt"

/*
	空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

接口值
一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。
*/

/*
想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

x.(T)
其中：

x：表示类型为interface{}的变量
T：表示断言x可能是的类型。
该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
*/

func main() {
	var x interface{}
	x = "鞠婧祎"
	res, ok := x.(string)
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println("是字符串类型", res)
	}
}

/*
上面的示例中如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现：

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
*/
