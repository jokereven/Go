package main;

import "fmt";

func main(){
	var a = []int{1,2,3,4}
	b := a;
	fmt.Println(a);	//[1,2,3,4]
	fmt.Println(b);	//[1,2,3,4]
	b[0] = 996;
	fmt.Println(a);	//[996 2 3 4]
	fmt.Println(b);	//[996 2 3 4]

	/*
	由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

	Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：

	copy(destSlice, srcSlice []T)
	其中：

	srcSlice: 数据来源切片
	destSlice: 目标切片
	*/

	//copy函数复制
	var n1 = []int{1,2,3,4}
	var n2 = []int{5,6,7,8}
	copy(n1,n2);
	//copy函数n1复制n2的
	fmt.Println(n1);	//[5 6 7 8]
	fmt.Println(n2);	//[5 6 7 8]

	n1[0] = 985;
	fmt.Println(n1);	//[985,6,7,8]
	fmt.Println(n2);	//[5,6,7,8]

	n2[0] = 985;
	fmt.Println(n1);	//[5,6,7,8]
	fmt.Println(n2);	//[985,6,7,8]
}
