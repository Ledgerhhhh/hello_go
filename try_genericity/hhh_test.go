package main

import (
	"fmt"
	"testing"
)

type myI1 interface {
	m1()
}

type myI2 interface {
	m2()
}

type interAll interface {
	myI1
	myI2
}

type mys1 struct {
}

func (m mys1) m1() {
	fmt.Println("m1")
}

type mys2 struct {
}

func (m mys2) m2() {
	fmt.Println("m2")
}

func hhhh[T interAll](i T) {
	fmt.Println(i)
}
func TestHHH(t *testing.T) {
	//i2 := mys2{}
	//hhhh(i2)
}
