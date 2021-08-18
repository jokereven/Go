package main

import "fmt"

func index(x int) {
	x = 520
}

func content(x *int) {
	*x = 520
}

func main() {
	a := 996
	index(a)
	fmt.Println(a) // 996
	content(&a)
	fmt.Println(a) // 520
}
