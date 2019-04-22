package main

import (
	"fmt"
)

//goroutine和channel的协同工作案例

func writeData(intChan chan int) {

	for i:= 1; i <= 50; i ++ {
		intChan <- i
		fmt.Println("writeData", i)
	}
	close(intChan)
}

func readChan(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Println("readData读取的数据为", v)
	}

	exitChan <- true
}

func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readChan(intChan, exitChan)

	for {
		_, ok := <- exitChan
		if !ok {
			break
		}
	}
}