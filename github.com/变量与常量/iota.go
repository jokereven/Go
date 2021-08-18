package main;

import "fmt";

// const (
// 	a1 = iota;	//0
// 	a2;
// 	a3;
// 	a4;
// )

// 使用_跳过某些值
// const (
// 	a1 = iota;	//0
// 	a2;	//1
// 	_;	//2	跳过
// 	a3;	//3
// )

// iota声明中间插队
const (
	a1 = iota;	//0
	a2 = 996;	//996
	a3 = iota;	//2
	a4;	//3 
)
const a5 = iota;
func main(){
	// fmt.Println(a1,a2,a3,a4);

	// fmt.Println(a1,a2,a3);

	fmt.Println(a1,a2,a3,a4,a5);
}
