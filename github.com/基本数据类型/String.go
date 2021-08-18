package main;

import (
	"fmt";
)

func main(){
	tosay := "我的名字叫卡卡西,我的愿望是成为火影";
	string := "这是字符串";

	fmt.Println(tosay);
	fmt.Println(string);

	// 字符串转义符

	/*
		转义符		含义
		\r			回车符（返回行首）
		\n			换行符（直接跳到下一行的同列位置）
		\t			制表符
		\'			单引号
		\"			双引号
		\\			反斜杠
	*/

	// 举个例子，我们要打印一个Windows平台下的一个文件路径：
 fmt.Println("str := \"d:\\Go\\gocode\\go.exe\"");

	//  多行字符串
	content :=`
	one
	two
	three
	`
	fmt.Println(content);

	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

	/*
		方法	介绍
		len(str)	求长度
		+或fmt.Sprintf	拼接字符串
		strings.Split	分割
		strings.contains	判断是否包含
		strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
		strings.Index(),strings.LastIndex()	子串出现的位置
		strings.Join(a[]string, sep string)	join操作
	*/

	// byte和rune类型

}
/*	Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符		*/
