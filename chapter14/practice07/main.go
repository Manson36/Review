package main

import (
	"encoding/json"
	"fmt"
)

//json练习

type Monster struct {
	Name string
	Age int
	Birthiday string
	Sal float64
	Skill string
}

func testStruct () {

	monster := Monster{
		Name:"牛魔王",
		Age:500,
		Birthiday:"2199-2-31",
		Sal:44444.4,
		Skill:"牛魔拳",

	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误 err", err)
	}
	fmt.Println("monster序列化后", string(data))
}

//将map进行序列化
func testMap () {

	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = "232"
	a["address"] = "火焰山"

	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Println("序列化错误 err", err)
	}
	fmt.Println("a map序列化后 ", string(data))
}

func main() {

	testStruct()
	testMap()

}