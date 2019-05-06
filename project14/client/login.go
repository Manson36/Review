package login

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/review/project14/common/message"
	"net"
	"time"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes的结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将LoginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//5.将data赋给mes.Data 字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//7.先把data的长度发送给服务器
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("conn.Write() err",err)
		return
	}
	fmt.Printf("客户端，发送消息的长度= %d，内容= %s", len(data), string(data))
	return

	//发送消息本身
	_, err = conn.Write(data)
	if err !=nil {
		fmt.Println("conn.Write() err",err)
		return
	}

	//休眠20s
	time.Sleep(20*time.Second)
	fmt.Println("休眠了20s")
	return
}