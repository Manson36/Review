package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	defer file.Close()
	//准备写入5句hello
	str := "hello"
	writer := bufio.NewWriter(file)
	for i:= 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//将在缓存的数据写入到文件之中
	writer.Flush()

}