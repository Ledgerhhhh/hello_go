package test

import (
	"fmt"
	"testing"
)

//各位请教个问题，如何实现两个结构体切片传入一个通用函数中，并且函数内能遍历切片并能取出每个切片的字段值

type Fields interface {
	GetAllField() []any
}
type myStruct struct {
	a string
	b int
	c float32
}

func (m myStruct) GetAllField() []any {
	return []any{m.a, m.b, m.c}
}

func TY(f Fields) {
	for _, v := range f.GetAllField() {
		fmt.Println(v)
	}
}

func TestHH(t *testing.T) {
	m := myStruct{
		a: "a",
		b: 1,
		c: 1.2,
	}
	TY(m)
}
