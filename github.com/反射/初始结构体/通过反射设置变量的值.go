package main

import (
	"fmt"
	"reflect"
)

/*
	想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，
	必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。

	想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
*/

/*
	func reflectchange1(x interface{}) {
		v := reflect.ValueOf(x)
		if v.Kind() == reflect.String {
			v.SetInt("鞠婧祎")	//修改的是副本 reflect包会引发panic
		}
	}
*/

func reflectchange2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(996)
	}
}

func main() {
	fmt.Println("")
	var n int = 717
	reflectchange2(&num)
	fmt.Println(n)
}
