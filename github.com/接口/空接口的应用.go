package main

import "fmt"

//空接口的应用
/*
	空接口作为函数的参数
	使用空接口实现可以接收任意类型的函数参数。
*/
/*
	空接口作为map的值
	使用空接口实现可以保存任意值的字典
*/

// func show(a interface{}) {
// 	fmt.Printf("type:%T value:%v\n", a, a)
// }

func main() {
	fmt.Println("应用")
	var Test = make(map[string]interface{}, 8)
	Test["name"] = "鞠婧祎"
	Test["age"] = 32
	Test["hobby"] = []string{"跳舞💃", "唱歌φ(*￣0￣)", "rep"}
	fmt.Println(Test)
}
