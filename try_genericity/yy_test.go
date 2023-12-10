package main

import (
	"fmt"
	"testing"
)

type myss1 struct {
}

func (m myss1) dealArray() {
	fmt.Println("hh")
}

type myss2 struct {
}

func (m myss2) dealArray() {
	fmt.Println("hh2")
}

type name interface {
	dealArray()
}

func d(n name) {
	n.dealArray()
}
func TestYyy(t *testing.T) {
	//m := myss1{}
	//d(m)
}
