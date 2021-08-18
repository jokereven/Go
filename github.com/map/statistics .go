package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "how do you do"
	var Map = make(map[string]int, 8)
	word := strings.Split(str, " ")
	for _, index := range word {
		val, ok := Map[index]
		if ok {
			Map[index] = val + 1
		} else {
			Map[index] = 1
		}
	}
	fmt.Println(Map)
}
