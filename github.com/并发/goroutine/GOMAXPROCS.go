package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果，这里举个例子：
*/

var wg sync.WaitGroup

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	wg.Done()
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(16)
	wg.Add(2)
	go a()
	go b()
	time.Sleep(time.Second)
	wg.Wait()
}

/*
	Go语言中的操作系统线程和goroutine的关系：

一个操作系统线程对应用户态多个goroutine。
go程序可以同时使用多个操作系统线程。
goroutine和OS线程是多对多的关系，即m:n。
*/
