package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("使用键值对初始化")

	a := Person{
		Name: "鞠婧祎",
		Age:  36,
	}
	b := Person{
		"杨超越",
		28,
	}
	fmt.Println(a, b)
}
