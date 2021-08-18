package main;

import "fmt";

func main(){
	/*
	切片的赋值拷贝:下面的代码中演示了拷贝前后两个变量共享底层数组，
	对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	*/

	a1 := []int{1,2,3,4,5,6}
	a2 := a1;
	a2[0] = 996;
	fmt.Println(a1);	//[996 2 3 4 5 6]
	fmt.Println(a2);	//[996 2 3 4 5 6]

	/*
	切片遍历:	切片的遍历方式和数组是一致的，支持索引遍历和for range遍历。
	*/
	var arr = []int{1,2,3,4,5,6,7}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//for range遍历
	for index,item := range arr{
		fmt.Println(index,item);
	}
	/*
	append()方法为切片添加元素
	Go语言的内建函数append()可以为切片动态添加元素。
	可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。
	*/
	var s []int	//[]
	s = append(s,1);	//[2]
	s = append(s, 2,3,4,5,6) //[1,2,3,4,5,6]
	var s1 = []int{996,985,211}
	s = append(s, s1...);
	fmt.Println(s);	//[1 2 3 4 5 6 996 985 211]
	/*
	注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	var s []int
	s = append(s, 1, 2, 3)
	*/
}
