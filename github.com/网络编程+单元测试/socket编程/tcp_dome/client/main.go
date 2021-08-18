package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
	一个TCP客户端进行TCP通信的流程如下：

建立与服务端的链接
进行数据收发
关闭链接
*/

func main() {
	//与服务端建立联系
	connect, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	//2. 利用该连接进行数据的发送和接受
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = connect.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := connect.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
