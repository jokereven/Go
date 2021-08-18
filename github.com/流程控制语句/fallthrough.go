package main

import "fmt"

//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

func switchDemo() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func main() {
	fmt.Println("fallthroungh")
	switchDemo() //a b
}
