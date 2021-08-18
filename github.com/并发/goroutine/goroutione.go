package main

import (
	"fmt"
	"sync"
)

/*
	goroutine

	在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。那么能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？

Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。
*/

/*

使用goroutine
Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。

一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。
*/

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello 鞠婧祎")
	wg.Done() //通知wg把计数器-1
}

func main() { //开启一个主goroutine去执行main函数
	wg.Add(1)  //计数牌🥧+1
	go hello() //开启了一个goroutine去执行hello函数
	fmt.Println("hello 杨超越")
	// time.Sleep(time.Second) //等待一秒
	wg.Wait() //等所有的小弟都干玩活才收兵
}
