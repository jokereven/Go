package main

import (
	"fmt";
)

/*
切片的长度和容量
切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
*/

/*
切片表达式
切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。
它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。
*/

func main(){
	// var arr = []string{"鞠婧祎","杨超越"}
	// fmt.Println(len(arr),cap(arr));
	// var a = []int{1,2,3,4,5,6}
	// s := a[1:3];
	// fmt.Printf("s:%v len(s):%v cap(s):%v",s,len(s),cap(s));	//s:[2 3] len(s):2 cap(s):5
	// /*
	// a[2:]  // 等同于 a[2:len(a)]
	// a[:3]  // 等同于 a[0:3]
	// a[:]   // 等同于 a[0:len(a)]
	// */
	// var index = []string{"卡卡西","鞠婧祎","杨超越","薛之谦","我爱罗"}
	// one := index[1:4];
	// fmt.Printf("one:%v len(one):%v cap(one):%v",one,len(one),cap(one));
	// two := one[1:3];
	// fmt.Printf("two:%v len(two):%v cap(two):%v",two,len(two),cap(two));
	// /*
	// 完整切片表达式
	// 对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：
	// a[low : high : max]
	// */
	var m = []int{1,2,3,4,5,6,7,8,9};
	n1 := m[4:6];	//从第四个往后数的算cap之前的不算
	fmt.Printf("n1:%v len(n1):%v cap(n1):%v",n1,len(n1),cap(n1));	//n1:[5 6] len(n1):2 cap(n1):5
	n2 := m[4:5:7];
	fmt.Printf("n2:%v len(n2):%v cap(n2):%v",n2,len(n2),cap(n2));	//n2:[5] len(n2):1 cap(n2):3
	// cap的值为max-high
}
