package main;

import "fmt";

func main() {
	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

/*
[0]  len:1  cap:1  ptr:0xc0000140a8
[0 1]  len:2  cap:2  ptr:0xc0000140f0
[0 1 2]  len:3  cap:4  ptr:0xc0000121c0
[0 1 2 3]  len:4  cap:4  ptr:0xc0000121c0
[0 1 2 3 4]  len:5  cap:8  ptr:0xc00000e2c0
[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc00000e2c0
[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc00000e2c0
[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc00000e2c0
[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc00007c080
[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc00007c080
*/

//append()函数将元素追加到切片的最后并返回该切片。
// 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
