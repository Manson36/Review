package main

import "fmt"

func main() {

	//演示管道使用
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的值%v，地址 %q\n", intChan, &intChan)

	intChan <- 10
	num := 200
	intChan <- num
	intChan <- 50

	fmt.Printf("channel len=%v, cap=%v\n", len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	fmt.Println("num2=",num2)
	fmt.Printf("channel len=%v, cap=%v\n", len(intChan), cap(intChan))

}