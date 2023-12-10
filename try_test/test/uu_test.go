package test

import (
	"fmt"
	"testing"
)

type mys11 struct {
	name string
}

type mys22 struct {
	name2 string
}

func hh[T mys11 | mys22](h []T) {
	for _, v := range h {
		fmt.Printf("%+v", v)
	}
}
func TestUU(t *testing.T) {
	m := []mys11{
		{
			name: "ledger",
		},
	}
	hh(m)
}
