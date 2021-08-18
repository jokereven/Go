package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁🔒

/*
	互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：
*/

//多个goroutine并发操作全局变量x

func add() {
	for i := 0; i < 500000; i++ {
		lock.Lock() //加🔒
		x = x + 1
		lock.Unlock() //解🔒
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

	使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
*/
