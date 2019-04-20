package main

import (
	"encoding/json"
	"fmt"
)

//反序列化实例
type Monster struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

func unmarshalStruct() {

	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthiday\":\"2199-2-31\",\"Sal\":44444.4,\"Skill\":\"牛魔拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarshal err", err)
	}
	fmt.Printf("反序列化后monster=%v，monster.Name=%v\n",monster, monster.Name)
	}

func unmarshalmap() {

	str := "{\"address\":\"火焰山\",\"age\":\"232\",\"name\":\"红孩儿\"}"

	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("unmarshal err", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}

func main() {

	unmarshalStruct()
	unmarshalmap()
}