package main

import (
	"fmt"
	"strings"
)

//字符串的一些常用的操作
func main()  {
	// // 长度 len()
	// var chinese string = "我的名字叫卡卡西";
	// fmt.Println(len(chinese));
	// var english string = "my name is kakaxi";
	// fmt.Println(len(english));

	// //拼接字符串
	// fmt.Println(chinese+english);

	// value := fmt.Sprintf("%s - %s",chinese, english);
	// fmt.Println(value);

	// //字符串分割
	// sli := "我的名字叫卡卡西";
	// fmt.Println(strings.Split(sli,""));

	// //判断是否包含
	// tains := "我的愿望是世界和平";
	// fmt.Println(strings.Contains(tains,"世界和平"));

	// 前缀/后缀判断
	has := "console.log";
	fmt.Println(strings.HasPrefix(has,"c"));	//前
	fmt.Println(strings.HasSuffix(has,"c"));	//后

	// 子串出现的位置
	index := "今天是个好日子";
	fmt.Println(strings.Index(index,"好"));
	fmt.Println(strings.LastIndex(index,"个"));

	// join操作
	src := [] string{"what", "do", "you", "main"};
	fmt.Println(src);
	fmt.Println(strings.Join(src,"-"));
}
