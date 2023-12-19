package main

import (
	"com.ledger.goproject/try_reflection/model"
	"fmt"
	"reflect"
	"testing"
)

func TestH(t *testing.T) {
	var x = 3

	reflect.TypeOf(x)
	reflect.ValueOf(x)
	//获取变量类型
	fmt.Println("type:", reflect.TypeOf(x))
	//获取变量的值
	fmt.Println("value:", reflect.ValueOf(x))
	//赋值
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(6.25)
	fmt.Println("Updated value of x:", x)
}

func TestB(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{
		Name: "ledger",
		Age:  10,
	}
	pType := reflect.TypeOf(p)
	pValue := reflect.ValueOf(p)
	for i := 0; i < pType.NumField(); i++ {
		// 类型信息
		field := pType.Field(i)
		// 值信息
		value := pValue.Field(i)
		// 和pType.Field(i)类似但是更为强大
		structField := pValue.Type().Field(i)

		fmt.Printf("FieldName:%s: === FieldType %v = FieldValue %v\n ", field.Name, field.Type, value)
		fmt.Printf("FieldName:%s: === FieldType %v = FieldValue %v\n ", structField.Name, structField.Type, value)
	}
}

func TestC(t *testing.T) {
	p2 := model.NewP("ledger", 20)
	v := reflect.ValueOf(p2).Elem()
	// 遍历结构体的字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		structField := v.Type().Field(i)
		fmt.Printf("FieldName:%s: === FieldType %v = FieldValue %v\n ", structField.Name, structField.Type, field)
	}
	// 根据名称获取字段的值
	name := v.FieldByName("Name")
	age := v.FieldByName("age")
	none := v.FieldByName("none")

	fmt.Println("Name:", name) //  Name: ledger
	fmt.Println("age:", age)   // age: 20
	fmt.Println("none:", none) // none: <invalid reflect.Value>
}

type p2 struct {
	name string
}

func (k *p2) Laugh() {
	fmt.Println("I am Laugh")
}

func TestD(t *testing.T) {
	p := &p2{}
	v := reflect.ValueOf(p)
	v.MethodByName("Laugh").Call(nil)
}

func TestE(t *testing.T) {
	p := model.NewP("ledger", 20)
	v := reflect.ValueOf(p)
	fmt.Println(v)
	PPMethod := v.MethodByName("PP")
	if PPMethod.IsValid() {
		PPMethod.Call(nil)
	} else {
		fmt.Printf("Method %s not found\n", PPMethod)
	}
	ppMethod := v.MethodByName("pp")
	if ppMethod.IsValid() {
		ppMethod.Call(nil)
	} else {
		fmt.Printf("Method %s not found\n", ppMethod)
	}
	noneMethod := v.MethodByName("none")
	if noneMethod.IsValid() {
		noneMethod.Call(nil)
	} else {
		fmt.Printf("Method %s not found\n", noneMethod)
	}
	sayMethod := v.MethodByName("Say")
	if sayMethod.IsValid() {
		call := sayMethod.Call([]reflect.Value{
			reflect.ValueOf("hello"),
		})
		fmt.Println("call", call[0].Interface().(string))
	} else {
		fmt.Printf("Method %s not found\n", sayMethod)
	}
}

func TestF(t *testing.T) {
	p := model.NewP("ledger", 20)
	t2 := reflect.ValueOf(p).Elem().Type()
	newP := reflect.New(t2)
	// 要设置的话需要先.Elem()获取指针,之后选择字段名名字之后set
	newP.Elem().FieldByName("Name").Set(reflect.ValueOf("jack"))
	// 设置不存在的字段会报错
	//newP.Elem().FieldByName("name").Set(reflect.ValueOf("jack"))
	// 设置不可访问的小写字段会报错
	//newP.Elem().FieldByName("age").Set(reflect.ValueOf(18))
	fmt.Println(newP.Interface().(*model.P))
}

func TestG(t *testing.T) {
	p := model.NewP("ledger", 20)
	v := reflect.ValueOf(p).Elem()
	v.FieldByName("Name").Set(reflect.ValueOf("jack"))
	//v.FieldByName("age").Set(reflect.ValueOf(18))
	fmt.Println("p", p)
}
func TestH2(t *testing.T) {
	p := model.NewP("ledger", 20)
	v := reflect.ValueOf(p).Elem()
	kind := v.FieldByName("hh").Type().Kind()
	fmt.Println(v.FieldByName("Name").Type())
	switch kind {
	case reflect.String:
		fmt.Println("Field is of type string")
	case reflect.Int:
		fmt.Println("Field is of type int")
	case reflect.Struct:
		fmt.Println("Field is of type struct")
	}
}

func TestK(t *testing.T) {
	newP2 := model.NewP2("ledger", 20)
	v := reflect.ValueOf(newP2).Elem()
	if v.Type().Field(1).Tag.Get("json") == "age" && v.Type().Field(1).Tag.Get("xml") == "age" {
		fmt.Println(v.Type().Field(1).Name)
		fmt.Println("ok")
	}
}

func TestU(t *testing.T) {
	p := model.NewP("ledger", 20)
	v := reflect.ValueOf(p).Elem()
	v.Field(0).Set(reflect.ValueOf("jack"))
	fmt.Println(*p)
}
