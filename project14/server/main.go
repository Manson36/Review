package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/review/project14/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn)(mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err", err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen])
	if n!= int(pkgLen) || err != nil {
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	return
}

//处理与客户端的通信
func process(conn net.Conn) {

	defer conn.Close()

	//循环的客户端发送信息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return
			}else {
				fmt.Println("readpkg err",err)
				return
			}
		}
		fmt.Println("mes=", mes)
	}
}

func main() {

	//提示信息
	fmt.Println("服务器在8889端口监听")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}

	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err", err)
		}

		//一旦连接成功，启用协程与客户端保持通讯
		go process(conn)
	}
}