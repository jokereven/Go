package main

import "fmt"

func main() {
	var joker_xue struct {
		name string
		age  int
	}
	joker_xue.name = "薛之谦"
	joker_xue.age = 36
	fmt.Printf("joker_xue === %v\n", joker_xue)  //joker_xue === {薛之谦 36}
	fmt.Printf("joker_xue === %#v\n", joker_xue) //joker_xue === struct { name string; age int }{name:"薛之谦", age:36}
}
