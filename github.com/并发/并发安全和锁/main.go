package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup

//多个goroutine并发操作全局变量x

func add() {
	for i := 0; i < 500000; i++ {
		x = x + 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	add()
	add()
	wg.Wait()
	fmt.Println(x)
}

/*
	上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。
*/
