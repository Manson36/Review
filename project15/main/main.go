package main

import (
	"fmt"
	"github.com/review/project08/model"
	"github.com/review/project08/service"
)

type customerView struct {
	key int
	loop bool

	customerService *service.CustomerService
}

func (this *customerView)exit() {

	fmt.Println("确认是否退出（y/n）：")
	choice := ""
	for {
		fmt.Scanln(&this.key)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}
	if choice == "y" {
		this.loop = false
	}

}

//得到用户输入的id，然后删除
func (this *customerView) delete() {

	fmt.Println("--------删除用户------")
	fmt.Println("请输入要删除的id")
	id := -1
	fmt.Println(&id)
	if id == -1 {
		return
	}
	fmt.Println("请确认是否删除（y/n）：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" {
		//调用service的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("--------删除完成--------")
		}else {
			fmt.Println("--------删除完成--------")
		}
	}
}

//得到用户的输入， 构建新用户，并完成添加
func (this *customerView)add() {

	fmt.Println("--------添加客户------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("email")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例，无id
	customer := model.NewCustomer2(name, gender, age, phone, email )
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("-------添加完成--------")
	} else {
		fmt.Println("-------添加失败--------")
	}
}

//显示所有客户所有信息
func (this *customerView)list() {

	//首先，获取当前所有的客户信息
	customers := this.customerService.List()

	//显示
	fmt.Println("----------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 1; i < len(customers); i ++ {

		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n----------客户列表完成----------\n")
}

//显示主菜单
func (this *customerView)mainMenu() {

	for {
		fmt.Println("------客户信息管理软件------")
		fmt.Println("        1.添加客户")
		fmt.Println("        2.修改客户")
		fmt.Println("        3.删除客户")
		fmt.Println("        4.客户列表")
		fmt.Println("        5.退    出")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&this.key)
		switch this.key {
		case 1:
			//fmt.Println("添加客户")
			this.add()
		case 2:
			fmt.Println("修改客户")
		case 3:
			fmt.Println("删除客户")
		case 4:
			//fmt.Println("客户列表")
			this.list()
		case 5:
			this.loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户管理系统")

}

func main() {

	customerView:= customerView{
		key:0,
		loop:true,
	}
	//这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}