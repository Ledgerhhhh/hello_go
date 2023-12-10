package main

import "fmt"

func main() {
	testT()
}

type mySlice[T int | float32 | float64] []T

func (s mySlice[T]) sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

type MyMap[KEY int | string, VALUE float32 | float64 | string] map[KEY]VALUE

type MyStruct[T int | string] struct {
	Name string
	Data T
}
type myInterface[T int | string] interface {
	print(data T)
}
type myChan[T int | string] chan T

type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

func Add[T int | float32 | float64 | string](a T, b T) T {
	return a + b
}

func MyFunc[T int | float32](a, b T) {
	// 匿名函数可使用已经定义好的类型形参
	fn2 := func(i T, j T) T {
		return i*2 - j*2
	}
	fn2(a, b)
}

type myStruct[T int | string] struct {
}

func (receiver myStruct[T]) add(a, b T) T {
	return a + b
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32
}

type Float interface {
	float32 | float64
}
type Slice[T Int | Uint | Float] []T // 使用 '|' 将多个接口类型组合

func testT() {
	//var ints mySlice[int] = []int{1, 2}
	//ints.sum()
	//var floats mySlice[float32] = []float32{1.0, 2.0}
	//var a MyMap[int, string] = map[int]string{
	//	1: "hhh",
	//	2: "hhh",
	//}
	//myStruct := MyStruct[int]{
	//	Name: "",
	//	Data: 1,
	//}
	//w := WowStruct[int, []int]{
	//	Data:     []int{1, 2},
	//	MaxValue: 0,
	//	MinValue: 0,
	//}
	add := Add("啊啊", "飒飒")
	fmt.Println(add)
}
