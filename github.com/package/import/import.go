package main

/*
导入包package

要在代码中引用其他包的内容，需要使用import关键字导入使用的包。具体语法如下:

import "包的路径"
注意事项：

import导入语句通常放在文件开头包声明语句的下面。
导入的包名需要使用双引号包裹起来。
包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
Go语言中禁止循环导入包。
单行导入
 */

import (
    "fmt"
    "/github.com/package/pck" //导入自定义的包
)

func main() {
    fmt.Println("这时外部导入的package")
    fmt.Println(Add(996,717))
}