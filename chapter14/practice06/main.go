package main

import (
"bufio"
"fmt"
"io"
"os"
)

//统计英文、数字、空格及其他字符的个数

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main() {

	fileName := "d:/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)
	//开始循环读取文件的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//为了兼容中文字符，将str转换成[]rune
		str2 := []rune(str)

		for _, v := range str2 {
			switch {
			case v >= 'a'&& v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\n':
				count.SpaceCount++
			case v >= 0 && v <= 9:
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数为%d，数字的个数为%d， 空格的个数为%d，其他字符的个数为%d",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
