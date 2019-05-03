package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp","192.168.20.253:8888")
	if err != nil {

		fmt.Println("client dial err",err)
		return
	}
	////功能一：客户端发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)

	//line, err := reader.ReadString('\n')
	//if err !=  nil {
	//
	//	fmt.Println("reader.ResdString err",err)
	//}
	////再将line发送给服务器
	//n, err := conn.Write([]byte(line))
	//if err != nil {
	//
	//	fmt.Println("conn.Write err ",err)
	//}
	//fmt.Printf("客户端发送了%d字节的数据，并退出", n)

	//改进
	 for {
		 line, err := reader.ReadString('\n')
		 if err !=  nil {

		 	fmt.Println("reader.ResdString err",err)
		 }

		 //如果用户输入的是exist就退出
		 line = strings.Trim(line, "\r\n")
		 if line == "exist" {
		 	fmt.Println("客户端退出")
		 	break
		 }

		 //再将line发送给服务器
		 _, err = conn.Write([]byte(line+"\n"))
		 if err != nil {
		 	fmt.Println("conn.Write err ",err)
		 }
	 }
}