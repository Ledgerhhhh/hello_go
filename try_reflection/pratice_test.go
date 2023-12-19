package main

import (
	"com.ledger.goproject/try_reflection/model"
	"fmt"
	"reflect"
	"testing"
)

var e = model.NewEmployee(1, "ledger", 20, 2000, "beijing")

func TestPre(t *testing.T) {
	//使用反射调用 AnnualSalary 方法，获取雇员的年薪。
	v := reflect.ValueOf(e)
	call := v.MethodByName("AnnualSalary").Call(nil)
	fmt.Println(call[0].Interface().(float64))
	//使用反射调用 UpdateLocation 方法，将雇员的位置更新为新的城市。
	values := v.MethodByName("UpdateLocation").Call([]reflect.Value{
		reflect.ValueOf("缅甸"),
	})
	fmt.Println(values)
	fmt.Printf("%+v", e)
	//使用反射调用 PrintInfo 方法，打印雇员的信息。
	v.MethodByName("PrintInfo").Call(nil)
	//使用反射调用 IncreaseSalaryByPercentage 方法，将雇员的薪水增加 10%。
	v.MethodByName("IncreaseSalaryByPercentage").Call([]reflect.Value{
		reflect.ValueOf(10.0),
	})
	fmt.Printf("%+v", e)
}

func TestPre2(t *testing.T) {
	////调用方法： 使用反射调用 AnnualSalary 方法，获取雇员的年薪。
	//v := reflect.ValueOf(e)
	//call := v.MethodByName("AnnualSalary").Call(nil)
	//fmt.Println(call[0].Interface().(float64))
	////标签解析： 创建一个通用的标签解析函数，输入一个结构体实例和一个标签名称，输出该标签对应字段的值。
	//value := common(e, "name")
	//fmt.Println(value.Interface())
	//类型判断： 编写一个函数，接受任意类型的参数，使用反射判断参数的类型，并执行相应的操作。考虑处理指针、数组、切片等多种情况。
	//exec(e)
	//修改字段值： 设计一个函数，接受一个结构体实例和一个 map[string]interface{}，根据 map 中的键值对，使用反射修改结构体对应字段的值。
	//m := map[string]any{}
	//m["Name"] = "kk"
	//editByMap(e, m)
	//fmt.Printf("%+v", e)
	//结构体初始化： 实现一个通用的结构体初始化函数，该函数接受结构体类型和一个 map[string]interface{}，返回一个用 map 中的值初始化的结构体实例。
	//data := map[string]interface{}{
	//	"ID":       101,
	//	"Name":     "Alice",
	//	"Age":      25,
	//	"Salary":   50000.0,
	//	"Location": "Wonderland",
	//}
	//editByMap(e, data)
	//fmt.Printf("%+v", e)
	//JSON 序列化： 编写一个通用的 JSON 序列化函数，该函数接受任意类型的结构体，使用反射将其转换为 JSON 格式的字符串。
	fmt.Printf("%+v", e)
	fmt.Println(structToJson(e))
	//标签验证： 创建一个函数，接受结构体实例和一个标签名称，检查结构体的每个字段是否都有该标签，如果没有，输出相应提示信息。

	//接口与反射结合： 设计一个接口，并实现多个结构体满足该接口。然后编写一个函数，接受实现了该接口的结构体实例，使用反射调用接口方法。

}

func structToJson(str any) string {
	var bytes []byte
	v := reflect.ValueOf(str)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ""
	}
	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		field := v.Field(i)
		name := structField.Name
		bytes = append(bytes, '"')
		bytes = append(bytes, []byte(name)...)
		bytes = append(bytes, '"')
		bytes = append(bytes, ':')
		// 处理字段的值
		switch field.Kind() {
		case reflect.String:
			bytes = append(bytes, '"')
			bytes = append(bytes, []byte(field.String())...)
			bytes = append(bytes, '"')
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			bytes = append(bytes, []byte(fmt.Sprintf("%d", field.Int()))...)
		case reflect.Float32, reflect.Float64:
			bytes = append(bytes, []byte(fmt.Sprintf("%f", field.Float()))...)
			// 添加其他可能的数据类型的处理
		case reflect.Struct:
			nestedJSON := structToJson(field.Interface())
			bytes = append(bytes, []byte(nestedJSON)...)
		case reflect.Array, reflect.Slice:
			bytes = append(bytes, '[')
			for j := 0; j < field.Len(); j++ {
				if j > 0 {
					bytes = append(bytes, ',')
				}
				nestedJSON := structToJson(field.Index(j).Interface())
				bytes = append(bytes, []byte(nestedJSON)...)
			}
			bytes = append(bytes, ']')
		}
		if i != v.NumField()-1 {
			bytes = append(bytes, ',')
			bytes = append(bytes, ' ')
		}
	}
	return "{" + string(bytes) + "}"
}

func initStruct(str any, m map[string]any) {
	for k, v := range m {
		setValue(str, k, v)
	}

}

func editByMap(str any, m map[string]any) {
	for k, v := range m {
		setValue(str, k, v)
	}
}
func setValue(str any, k string, v any) {
	val := reflect.ValueOf(str)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if field.Name == k {
			val.Field(i).Set(reflect.ValueOf(v))
		}
	}
}

func exec(str any) {
	v := reflect.ValueOf(str)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switchField(field)
	}
}

func switchField(v reflect.Value) {
	switch v.Kind() {
	case reflect.String:
		fmt.Println("string type", v.String())
	case reflect.Int:
		fmt.Println("int type", v.Int())
	case reflect.Bool:
		fmt.Println("bool type", v.Bool())
	case reflect.Ptr:
		fmt.Println("ptr type")
		exec(v.Interface())
	case reflect.Struct:
		fmt.Println("struct type")
		exec(v.Interface())
	case reflect.Slice:
		fmt.Println("slice type", v.Slice(0, 1).Interface())
		// 遍历切片信息
		for i := 0; i < v.Len(); i++ {
			switchField(v.Index(i))
		}
	case reflect.Array:
		fmt.Println("array type", v.Slice(0, 2).Interface())
		for i := 0; i < v.Len(); i++ {
			switchField(v.Index(i))
		}
	default:
		fmt.Println("other type")
	}
}
func common(str any, tagName string) reflect.Value {
	v := reflect.ValueOf(str)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		get := v.Type().Field(i).Tag.Get("json")
		if get == tagName {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}
