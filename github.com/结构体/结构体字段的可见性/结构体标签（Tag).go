package main

import (
	"encoding/json"
	"fmt"
)

/*
	Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：

`key1:"value1" key2:"value2"`
结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。

注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

例如我们为Student结构体的每个字段定义json序列化时使用的Tag：
*/

type Student struct {
	Name string `json:"Id"` //通过指定tag实现json序列化该字段时的key
	Age  int    //json序列化是默认使用字段名作为key
	love string //私有无法被json包访问
}

func main() {

	s1 := Student{
		Name: "鞠婧祎",
		Age:  36,
		love: "跳舞",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("failed")
		return
	} else {
		fmt.Printf("json data:%s\n", data)
	}
	fmt.Println("结构体标签（Tag）") //son data:{"Id":"鞠婧祎","Age":36}
}
