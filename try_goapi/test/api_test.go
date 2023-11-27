package test

import (
	"fmt"
	"testing"
)

func PrintArr[T string | int](arr []T) {
	for i := range arr {
		fmt.Println(i)
	}
}

// 所有类型
func PrintArr2[T any](arr []T) {
	for i := range arr {
		fmt.Println(i)
	}
}

// 包括但不限于整数、浮点数、字符串
func PrintArr3[T comparable](arr []T) {
	for i := range arr {
		fmt.Println(i)
	}
}

func TestT(t *testing.T) {
	PrintArr([]int{1, 2, 3, 4})
}
