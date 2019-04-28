package model

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//使用工厂模式返回一个Customer的实例
func NewCustomer (id int, name string, gender string, age int, phone string , email string)Customer {
	return Customer{
		id,
		name,
		gender,
		age,
		phone,
		email,
	}
}