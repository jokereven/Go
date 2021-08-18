package main

import "fmt"

//定义具有多放回值的函数
func ret(a, b int) (sum, sun int) {
	sum = a + b
	sun = a - b
	return
}
func main() {
	fmt.Println("return")

	sum, sun := ret(1314, 520)
	fmt.Println(sum, sun)
}
