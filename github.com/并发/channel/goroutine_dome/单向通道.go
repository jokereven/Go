package main

import "fmt"

/*
	两个goroutine,两个channel
		1. 生成0~100的数字🔢发送到ch1中
		2. 从ch1中取出数据计算它的平方,把结果发送到ch2中
*/

//生成0~100的数字🔢发送到ch1中
func f1(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	//从通道中取值的方法一
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		} else {
			ch2 <- tmp * tmp
		}
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	//for range取值
	for ret := range ch2 {
		fmt.Println(ret)
	}
}

/*
	有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。
*/

/*
	其中，

chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
<-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
在函数传参及任何赋值操作中可以将双向通道转换为单向通道，但反过来是不可以的。
*/
