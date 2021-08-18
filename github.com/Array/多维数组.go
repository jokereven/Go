package main;

import "fmt";

//多维数组
func main(){
	arr := [2][2]string{
		{"鞠婧祎","薛之谦"},
		{"毛不易","卡卡西"},
	}
	fmt.Println(arr);
	fmt.Println(arr[0][0]);

	// 二维数组的遍历
	//for❄循环遍历
	for i := 0; i < len(arr[i]); i++ {
		for j := 0; j <= i; j++{
			fmt.Println(arr[i][j])
		}
	}

	// for❄range遍历
	for _, one := range arr {
		for _, two := range one {
			fmt.Printf("%s\t",two);
		}
		fmt.Println();
	}

	// 注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：

	// 支持的写法
	{/*a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}

	//不支持多维数组的内层使用...
	b := [3][...]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	*/}
}
