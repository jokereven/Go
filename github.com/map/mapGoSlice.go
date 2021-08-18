package main

import (
	"fmt"
)

func main() {
	/*
		元素为map类型的切片
	*/
	var MapSlice = make([]map[string]string, 4)
	// 对切片中的map元素进行初始化
	MapSlice[0] = make(map[string]string, 3)
	MapSlice[0]["username"] = "卡卡西"
	MapSlice[0]["passwd"] = "123456"
	for index, value := range MapSlice[0] {
		fmt.Println(index, value) //passwd 123456		username 卡卡西
	}

	/*
		值为切片类型的map
	*/
	var SclieMap = make(map[string][]int, 3)
	//对切片初始化
	SclieMap["卡卡西"] = make([]int, 3)
	SclieMap["卡卡西"][1] = 117
	SclieMap["卡卡西"][2] = 317
	for key, value := range SclieMap {
		fmt.Println(key, value)
	}
}
