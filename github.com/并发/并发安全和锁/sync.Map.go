package main

import (
	"fmt"
	"strconv"
	"sync"
)

// //Go语言中内置的map不是高并发安全的。

// var (
// 	wg sync.WaitGroup
// )

// var m = make(map[int]int)

// func get(key int) int {
// 	return m[key]
// }

// func set(key int, value int) {
// 	m[key] = value
// }

// func main() {
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			set(i, i+100)                            //设置值 i key i+100 value
// 			fmt.Printf("key:%v value:%v", i, get(i)) //取值
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

var m = sync.Map{} //sync.map

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
