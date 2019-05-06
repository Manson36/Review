package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {

	//接受用户的选择
	var key int
	var loop bool = true

	for loop {
		fmt.Println("--------欢迎登陆多人聊天系统------")
		fmt.Println("\t 1.登陆聊天室")
		fmt.Println("\t 2.注册用户")
		fmt.Println("\t 3.退出系统")
		fmt.Println("请选择（1-3）")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登陆多人聊天系统")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	if key == 1 {
		//说明要登陆系统
		fmt.Println("请输入用户的id")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanln(&userPwd)

		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登陆失败")
		}else {
			fmt.Println("登陆成功")
		}
	}else if key == 2 {
		fmt.Println("用户注册的逻辑")
	}
}