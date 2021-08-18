package main

import "fmt"

/*
匿名导入包
如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。具体的格式如下：

import _ "包的路径"
匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。
 */

func main() {
    fmt.Println("匿名import")
}