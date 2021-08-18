package main

import "fmt"

func main() {
	/*
		map的定义

		Go语言中 map的定义语法如下：

		map[KeyType]ValueType其中，

		KeyType:表示键的类型。
		ValueType:表示键对应的值的类型。
		map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

		make(map[KeyType]ValueType, [cap])
		其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
	*/

	Map := make(map[string]int, 4)
	Map["鞠婧祎"] = 24
	Map["杨超越"] = 22
	fmt.Println(Map)                  //map[杨超越:22 鞠婧祎:24]
	fmt.Println(Map["鞠婧祎"])           //24
	fmt.Printf("Map of type:%T", Map) //Map of type:map[string]int

	// map也支持在声明的时候填充元素，例如：
	Userinfo := map[string]string{
		"username": "卡卡西",
		"passwd":   "123456",
	}
	fmt.Println(Userinfo) //Map of type:map[string]intmap[passwd:123456 username:卡卡西]

	/*
		判断某个键是否存在
		Go语言中有个判断map中键是否存在的特殊写法，格式如下:
		value, ok := map[key]
	*/

	Exist := make(map[string]int)
	Exist["卡卡西"] = 30
	Exist["我爱罗"] = 28
	value, ok := Exist["卡卡西"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不在这")
	}

	/*
		map的遍历
		Go语言中使用for range遍历map。
	*/

	For := make(map[string]int)
	For["卡卡西"] = 30
	For["自来也"] = 42
	For["鞠婧祎"] = 26
	for key, value := range For {
		fmt.Println(key, value)
	}
	//但我们只想遍历key的时候，可以按下面的写法：
	for key := range For {
		fmt.Println(key)
	}

	/*
		使用delete()函数删除键值对
		使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：

		delete(map, key)
		其中，

		map:表示要删除键值对的map
		key:表示要删除的键值对的键
		示例代码如下：
	*/

	Del := make(map[string]int)
	Del["鞠婧祎"] = 26
	Del["卡卡西"] = 30
	Del["我爱罗"] = 28
	fmt.Println(Del)
	delete(Del, "卡卡西")
	fmt.Println(Del)
}
