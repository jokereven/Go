package main;

import "fmt";

/*切片（Slice）是一个拥有相同类型元素的可变长度的序列。
它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
切片是一个引用类型，它的内部结构包含地址、长度和容量。
切片一般用于快速地操作一块数据集合。*/

func main(){
	// 声明切片类型的基本语法如下：var name []T	name:表示变量名	T:表示切片中的元素类型
	var a []string;	//生明一个字符串切片
	var b = []int{};	//声明一个整型切片并初始化
	var c = []bool{false,true}	//声明一个布尔切片并初始化
	fmt.Println(a);	//[]
	fmt.Println(b);	//[]
	fmt.Println(c);	//[false true]
	fmt.Printf("%T",a); //[]string
	fmt.Printf("%T",b); //[]int
	fmt.Printf("%T",c); //[]bool
	if (a==nil/*nil代表无,即声明未定义*/) {
		fmt.Println("nil");
	}
	if (b==nil) {
		fmt.Println("nil");
	}
	if (c==nil) {
		fmt.Println("nil");
	}
}
