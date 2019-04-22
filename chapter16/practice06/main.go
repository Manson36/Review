package main

import (
	"fmt"
	"time"
)

//统计1-8000中哪些数是素数

//向intChan中放入1-8000
func putNum(intChan chan int) {

	for i := 1; i <= 8000; i++{
		intChan <- i
	}
	close(intChan)
	}

//从intChan中取出数据，并判断是否是素数
//如果是放入到primeChan中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum因为取不到数，退出")
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)
	//开启四个协程，从intChan中取出数据判断是否是素数
	for i:= 0; i < 4; i ++ {
		go primeNum(intChan , primeChan, exitChan)
	}

	//这里我们主线程，进行处理
	go func() {
		for i := 0; i <4; i ++ {
			<-exitChan
		}
		close(primeChan)
	}()

	//遍历我们的primeChan，将结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数有", res)
	}
	fmt.Println("main线程退出")
}