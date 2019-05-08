package service

import "github.com/review/project08/model"

type CustomerService struct {
	customers []model.Customer
	customerNum int
}

func NewCustomerService () *CustomerService {

	//为了能在切片中看到客户，我们先初始化一个用户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三", "男", 20, "112", "js@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this CustomerService)List() []model.Customer{

	return this.customers
}

//添加客户到customers切片
func (this *CustomerService)Add(customer model.Customer)bool {

	//我们确定一个添加id的规则
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers)
	return true
}

func (this *CustomerService)FindById (id int) int {

	index := -1
	//遍历customers的切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//找到
			index = i
		}
	}
	return index
}

func (this *CustomerService)Delete(id int) bool {

	index := this.FindById(id)

	if id == -1 {
		return false
	}
	//如何删除切片中的一个元素
	this.customers = append(this.customers[:index], this.customers[index + 1:]...)
	return true
}
