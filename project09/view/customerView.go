package view

import (
	"fmt"
	"github.com/review/project08/model"
	"github.com/review/project09/service"
)

type customerView struct {
	key string
	loop bool

	customerService *service.CustomerService
}

func (this *customerView)delete() {

	fmt.Println("-------删除客户---------")
	fmt.Println("请选择删除客户的id：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("请确认是否删除（y/n）")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("---------删除完成--------")
		}else {
			fmt.Println("---------删除失败--------")
		}
	}
}

func (this customerView)add() {

	fmt.Println("---------添加客户--------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&name)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("--------添加完成---------")
	}else {
		fmt.Println("--------添加失败---------")
	}

}

func (this *customerView)list() {

	customers := this.customerService.List()

	fmt.Println("---------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("---------客户列表完成----------")
}

//显示主菜单
func (this *customerView)mainMenu() {

	for {
		fmt.Println("--------客户信息管理软件---------")
		fmt.Println("           1.添加客户")
		fmt.Println("           2.修改客户")
		fmt.Println("           3.删除客户")
		fmt.Println("           4.客户列表")
		fmt.Println("           5.退出")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			//fmt.Println("添加客户")
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			//fmt.Println("删除客户")
			this.delete()
		case "4":
			//fmt.Println("客户列表")
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户管理软件")
}

func main()  {

	customerView := customerView{
		key:"",
		loop:true,
	}

	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}