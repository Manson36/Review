package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	//这里我们循环接受客户端发送的消息
	defer conn.Close()

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待客户端 %s发送消息\n",conn.RemoteAddr().String())

		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("客户端退出 err=", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {

		fmt.Println("listen err =", err)
		return
	}
	defer listen.Close()

	//循环等待客户端来连接
	fmt.Println("等待客户端连接")
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("listen.Accept err", err)

	}else {
		fmt.Printf("Accept() sus conn=%v 客户端IP= %v\n", conn, conn.RemoteAddr().String())

	}
	//这里准备一个协程，为客户端服务
	go process(conn)
}