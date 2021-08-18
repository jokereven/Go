package main

import "fmt"

func main() {
	var ch1 chan int        //引用类型需要初始化之后才可以使用
	ch1 = make(chan int)    //初始化 无缓冲的通道 又称为阻塞的通道 同步通道
	ch1 = make(chan int, 1) //有缓冲的通道 异步通道
	ch1 <- 10               //把10发送到ch1中
	x := <-ch1              //取值

	//我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
	// len(ch1)
	// cap(ch1)

	fmt.Println(x)
}
