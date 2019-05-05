package main

import "fmt"

func login(userId int, userPwd string) (err error) {

	fmt.Printf("userId = %v, userPwd = %v\n", userId, userPwd)

	return nil
}