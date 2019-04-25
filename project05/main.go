package main

import "fmt"

type FamilyAccount struct {
	key int
	loop bool
	balance float64
	money float64
	note string
	details string
	flag bool
}

func NewFamilyAccount () *FamilyAccount {

	return &FamilyAccount{
		key:0,
		loop:true,
		balance:10000.0,
		money:0.0,
		note:"",
		details:"收支\t账户余额\t收支金额\t说  明",
		flag:true,
	}
}

func (this *FamilyAccount)showDetails() {

	fmt.Println("------当前记账收支明细------")
	if this.flag {
		fmt.Println("当前没有明细，来一笔吧")
	}else {
		fmt.Println(this.details)
	}
}

func (this *FamilyAccount)incomes() {

	fmt.Println("本次收支金额")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入的说明")
	fmt.Scanln(&this.note)

	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v\t", this.balance, this.money, this.note)
	this.flag = false
}

func (this *FamilyAccount)pay () {

	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出的说明")
	fmt.Scanln(&this.note)

	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v\t", this.balance, this.money, this.note)
	this.flag = false
}

func (this *FamilyAccount)exit() {

	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("您的输入有误，请重新输入")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount)MainMenu() {

	for {
		fmt.Println("-----家庭收支记账软件------")
		fmt.Println("        1.收支明细")
		fmt.Println("        2.登记收入")
		fmt.Println("        3.登记支出")
		fmt.Println("        4.退出软件")

		fmt.Println("请选择（1-4）")
		fmt.Scanln(&this.key)

		switch this.key {
		case 1:
			this.showDetails()
		case 2:
			this.incomes()
		case 3:
			this.pay()
		case 4:
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
}

func main() {

	NewFamilyAccount().MainMenu()

}