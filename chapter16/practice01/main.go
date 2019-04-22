package main

import (
	"fmt"
	"strconv"
	"time"
)

func Test() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Test() Hello World" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go Test()

	for i := 1; i <= 10; i++ {
		fmt.Println("Main() Hello World" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}