package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
*/

//读写互斥锁

var x int
var wg sync.WaitGroup
// var lock sync.Mutex //互斥锁🔒 15.9467248s
var rwlock sync.RWMutex

func read() {
	// lock.Lock()
	rwlock.Lock()
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}

func write() {
	// lock.Lock()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	// lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now() //当前的时间
	for i := 0; i < 996; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
