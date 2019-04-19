package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("d:/BugReport.txt")
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("file = %v",file)
	err = file.Close()
	if err != nil {
		fmt.Println("file.Close() err", err)
	}
}