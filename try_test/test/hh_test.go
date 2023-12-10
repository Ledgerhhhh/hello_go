package test

import (
	"fmt"
	"testing"
)

type mys1 struct {
	name string
}

type mys2 struct {
	name2 string
}

type fff interface {
	mys1 | mys2
}

func GetAllFiled[T fff](hh []T) {
	for i := range hh {
		fmt.Println(hh[i])
	}
}
func TestG(t *testing.T) {
	m := []mys1{
		{
			name: "ledger",
		},
	}
	GetAllFiled[mys1](m)
}
