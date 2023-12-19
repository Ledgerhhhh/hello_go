package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 3

	//获取变量类型
	fmt.Println("type:", reflect.TypeOf(x))
	//获取变量的值
	fmt.Println("value:", reflect.ValueOf(x))
	//赋值
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(6.25)
	fmt.Println("Updated value of x:", x)
}
