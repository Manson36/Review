package main

import (
	"fmt"
	"time"
)

//需求：计算1-200的阶乘，并把各个数的阶乘放到map中

//思路：
//1.编写函数计算各个数的阶乘，并放到map中
//2.我们启动协程多个，将统计的结果放入到map中
//3.map应该做成一个全局的

var myMap = make(map[int]int, 10)

func test(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	myMap[n] = res
}


func main() {

	//这里我们开启多个协程完成多个任务[200]
	for i := 1; i <= 200; i++ {
	go test(i)
}

	time.Sleep(time.Second *10)

	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}