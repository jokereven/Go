package main

import "fmt"

/*
	两个goroutine,两个channel
		1. 生成0~100的数字🔢发送到ch1中
		2. 从ch1中取出数据计算它的平方,把结果发送到ch2中
*/

//生成0~100的数字🔢发送到ch1中
func f1(ch1 chan int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 chan int, ch2 chan int) {
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
