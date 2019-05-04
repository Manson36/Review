package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//发射的最佳实践

type Monster struct {
	Name string `json:"monster_name"`
	Age int
	Score float64
	Sex string
}

//方法：显示s的值
func(s Monster)Print () {
	fmt.Println("-------Start------")
	fmt.Println(s)
	fmt.Println("-------End------")
}

func TextStruct(a interface{}) {

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd :=val.Kind()

	if kd != reflect.Ptr && val.Elem().Kind()  == reflect.Struct {
		fmt.Println("except struct")
		return
	}
	num := val.Elem().NumMethod()
	val.Elem().Field(0).SetString("白象精")

	//遍历结构体的所有字段
	for i := 0; i<num; i++ {

		fmt.Printf("Field%d:的值为%v\n", i, val.Elem().Field(i))
	}
	fmt.Printf("struct has %d feilds\n", num)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag = %s \n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d feilds\n", numOfMethod)

	val.Elem().Method(0).Call(nil)
}

func main() {

	var a Monster = Monster{
		Name:"黄狮子",
		Age:444,
		Score:3333.33,
	}

	result, _ := json.Marshal(a)
	fmt.Println("json Marshal", string(result))

	TextStruct(&a)
	fmt.Println(a)
}