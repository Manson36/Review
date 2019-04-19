package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//拷贝文件

func CopyFile(dstFileName string, srcFileName string)(written int64, err error){

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err", err)

	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {

	srcFile := "d:/qq.jpg"
	dstFile := "d:/aa.txt"
	_, err := CopyFile(dstFile, srcFile)

	if err != nil {
		fmt.Println("拷贝失败")
	}
	fmt.Println("拷贝完成")
}