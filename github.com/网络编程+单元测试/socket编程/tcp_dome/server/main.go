package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(connect net.Conn) {
	defer connect.Close() //处理完成之后关闭连接
	//争对当前获取到的数据进行发送和接收的操作
	for {
		reader := bufio.NewReader(connect)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		connect.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	/*
		1. 监听端口
		2. 接收客户端请求建立链接
		3. 创建goroutine处理链接。
	*/

	//监听端口
	listen, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//等待客户端来连接
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 启动一个单独的goroutine去处理连接
		go process(connect)
	}
}
