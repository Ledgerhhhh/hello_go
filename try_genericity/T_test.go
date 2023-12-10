package main

import (
	"fmt"
	"testing"
)

type s1 struct {
	Name1 string
}

type s2 struct {
	Name2 string
}

type my interface {
	s1 | s2
}

func hh[T my](t []T) {
	for _, v := range t {
		fmt.Println(v)
	}
}

func Test(t *testing.T) {
	var s1s []s1
	s1s = append(s1s, struct{ Name1 string }{Name1: "232"})
	hh(s1s)

	var s2s []s2
	s2s = append(s2s, struct{ Name2 string }{Name2: "1212125"})
	hh(s2s)
}
