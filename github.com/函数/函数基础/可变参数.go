package main

import "fmt"

//可变参数类型为切片
//固定参数和可变参数同时出现可变参数要放在后面
//Go语言没有默认参数
func intSum(a ...int) int {
	fmt.Println(a)      //[520 1314]
	fmt.Printf("%T", a) //[]int

	ret := 0
	for _, index := range a {
		ret += index
	}
	return ret
}

func main() {
	fmt.Println("Hello World")

	intSum(520, 1314)

	con := intSum(1314, 520)
	fmt.Println(con)
}
