package main;

import "fmt";

/*
使用make()函数构造切片
我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：

make([]T, size, cap)
其中：

T:切片的元素类型
size:切片中元素的数量
cap:切片的容量

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

判断切片是否为空
要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
*/


func main()  {
	arr := make([]int, 4, 6);
	fmt.Printf("arr:%v len(arr):%v cap(arr):%v",arr,len(arr),cap(arr));	//arr:[0 0 0 0] len(arr):4 cap(arr):6
	if (len(arr)==0) {
		fmt.Println("null");
	}
}

/*
切片不能直接比较
切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例：

var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
*/
