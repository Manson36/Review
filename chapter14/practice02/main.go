package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("d:/BugReport.txt")
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("读取文件结束")
}